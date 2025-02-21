package filesystem

import (
	"testing"

	"github.com/prysmaticlabs/prysm/v4/config/params"
	"github.com/spf13/afero"
)

// NewEphemeralBlobStorage should only be used for tests.
// The instance of BlobStorage returned is backed by an in-memory virtual filesystem,
// improving test performance and simplifying cleanup.
func NewEphemeralBlobStorage(_ testing.TB) *BlobStorage {
	return &BlobStorage{fs: afero.NewMemMapFs()}
}

// NewEphemeralBlobStorageWithFs can be used by tests that want access to the virtual filesystem
// in order to interact with it outside the parameters of the BlobStorage api.
func NewEphemeralBlobStorageWithFs(_ testing.TB) (afero.Fs, *BlobStorage, error) {
	fs := afero.NewMemMapFs()
	retentionEpoch := params.BeaconNetworkConfig().MinEpochsForBlobsSidecarsRequest
	return fs, &BlobStorage{fs: fs, retentionEpochs: retentionEpoch}, nil
}

type BlobMocker struct {
	fs afero.Fs
	bs *BlobStorage
}

// CreateFakeIndices creates empty blob sidecar files at the expected path for the given
// root and indices to influence the result of Indices().
func (bm *BlobMocker) CreateFakeIndices(root [32]byte, indices []uint64) error {
	for i := range indices {
		n := blobNamer{root: root, index: indices[i]}
		if err := bm.fs.MkdirAll(n.dir(), directoryPermissions); err != nil {
			return err
		}
		f, err := bm.fs.Create(n.path())
		if err != nil {
			return err
		}
		if err := f.Close(); err != nil {
			return err
		}
	}
	return nil
}

// NewEphemeralBlobStorageWithMocker returns a *BlobMocker value in addition to the BlobStorage value.
// BlockMocker encapsulates things blob path construction to avoid leaking implementation details.
func NewEphemeralBlobStorageWithMocker(_ testing.TB) (*BlobMocker, *BlobStorage) {
	fs := afero.NewMemMapFs()
	bs := &BlobStorage{fs: fs}
	return &BlobMocker{fs: fs, bs: bs}, bs
}
