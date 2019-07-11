package certwatcher

import (
	"context"
	"crypto/tls"
	"sync"

	"github.com/solo-io/go-utils/contextutils"
	"gopkg.in/fsnotify.v1"
)

// CertWatcher watches certificate and key files for changes.  When either file
// changes, it reads and parses both and calls an optional callback with the new
// certificate.

type certWatcher struct {
	sync.RWMutex

	currentCert *tls.Certificate
	watcher     *fsnotify.Watcher

	certPath string
	keyPath  string
}

// New returns a new certWatcher watching the given certificate and key.
func New(ctx context.Context, certPath, keyPath string) (*certWatcher, error) {
	var err error

	cw := &certWatcher{
		certPath: certPath,
		keyPath:  keyPath,
	}

	// Initial read of certificate and key.
	if err := cw.ReadCertificate(ctx); err != nil {
		return nil, err
	}

	cw.watcher, err = fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	return cw, nil
}

// GetCertificate fetches the currently loaded certificate, which may be nil.
func (cw *certWatcher) GetCertificate(_ *tls.ClientHelloInfo) (*tls.Certificate, error) {
	cw.RLock()
	defer cw.RUnlock()
	return cw.currentCert, nil
}

// Start starts the watch on the certificate and key files.
func (cw *certWatcher) Start(ctx context.Context) error {
	files := []string{cw.certPath, cw.keyPath}

	for _, f := range files {
		if err := cw.watcher.Add(f); err != nil {
			return err
		}
	}

	contextutils.LoggerFrom(ctx).Debug("Starting certificate watcher")
	cw.Watch(ctx)

	return cw.watcher.Close()
}

// Watch reads events from the watcher's channel and reacts to changes.
func (cw *certWatcher) Watch(ctx context.Context) {
	for {
		select {
		case event, ok := <-cw.watcher.Events:
			// Channel is closed.
			if !ok {
				return
			}

			cw.handleEvent(ctx, event)

		case err, ok := <-cw.watcher.Errors:
			// Channel is closed.
			if !ok {
				return
			}

			contextutils.LoggerFrom(ctx).Error(err, "certificate watch error")
		case <-ctx.Done():
			contextutils.LoggerFrom(ctx).Info("context closed")
			return
		}

	}
}

// ReadCertificate reads the certificate and key files from disk, parses them,
// and updates the current certificate on the watcher.  If a callback is set, it
// is invoked with the new certificate.
func (cw *certWatcher) ReadCertificate(ctx context.Context) error {
	cert, err := tls.LoadX509KeyPair(cw.certPath, cw.keyPath)
	if err != nil {
		return err
	}

	cw.Lock()
	defer cw.Unlock()
	cw.currentCert = &cert

	contextutils.LoggerFrom(ctx).Info("Updated current TLS certiface")

	return nil
}

func (cw *certWatcher) handleEvent(ctx context.Context, event fsnotify.Event) {
	// Only care about events which may modify the contents of the file.
	if !(isWrite(event) || isRemove(event) || isCreate(event)) {
		return
	}

	contextutils.LoggerFrom(ctx).Info("certificate event", "event", event)

	// If the file was removed, re-add the watch.
	if isRemove(event) {
		if err := cw.watcher.Add(event.Name); err != nil {
			contextutils.LoggerFrom(ctx).Error(err, "error re-watching file")
		}
	}

	if err := cw.ReadCertificate(ctx); err != nil {
		contextutils.LoggerFrom(ctx).Error(err, "error re-reading certificate")
	}
}

func isWrite(event fsnotify.Event) bool {
	return event.Op&fsnotify.Write == fsnotify.Write
}

func isCreate(event fsnotify.Event) bool {
	return event.Op&fsnotify.Create == fsnotify.Create
}

func isRemove(event fsnotify.Event) bool {
	return event.Op&fsnotify.Remove == fsnotify.Remove
}