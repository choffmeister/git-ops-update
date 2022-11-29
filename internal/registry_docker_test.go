package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDockerFetchVersions(t *testing.T) {
	reg1 := DockerRegistry{
		Url: "https://ghcr.io",
	}
	output1, err := reg1.FetchVersions("airfocusio/git-ops-update")
	assert.NoError(t, err)
	assert.Greater(t, len(output1), 0)

	reg2 := DockerRegistry{
		Url: "https://quay.io",
	}
	output2, err := reg2.FetchVersions("oauth2-proxy/oauth2-proxy")
	assert.NoError(t, err)
	assert.Greater(t, len(output2), 0)
}

func TestDockerRetrieveMetadata(t *testing.T) {
	reg1 := DockerRegistry{
		Url: "https://ghcr.io",
	}
	output1a, err := reg1.RetrieveMetadata("airfocusio/git-ops-update-test", "docker-v2-manifest-list-v0.0.1")
	assert.NoError(t, err)
	assert.Equal(t, map[string]string{
		"io.buildah.version":                "1.23.1",
		"org.opencontainers.image.version":  "0.0.1",
		"org.opencontainers.image.source":   "https://github.com/airfocusio/git-ops-update",
		"org.opencontainers.image.revision": "99a7aaa2eff5aff7c77d9caf0901454fedf6bf00",
	}, output1a)

	output1b, err := reg1.RetrieveMetadata("airfocusio/git-ops-update-test", "oci-v1-image-index-v0.0.1")
	assert.NoError(t, err)
	assert.Equal(t, map[string]string{
		"io.buildah.version":                "1.23.1",
		"org.opencontainers.image.version":  "0.0.1",
		"org.opencontainers.image.source":   "https://github.com/airfocusio/git-ops-update",
		"org.opencontainers.image.revision": "99a7aaa2eff5aff7c77d9caf0901454fedf6bf00",
	}, output1b)

	reg2 := DockerRegistry{
		Url: "https://quay.io",
	}
	output2, err := reg2.RetrieveMetadata("oauth2-proxy/oauth2-proxy", "v7.3.0")
	assert.NoError(t, err)
	assert.Equal(t, map[string]string{}, output2)
}

func TestDockerRetrieveDockerV2ManifestMetadata(t *testing.T) {
	reg := DockerRegistry{
		Url: "https://ghcr.io",
	}
	output, err := reg.RetrieveMetadata("airfocusio/git-ops-update-test", "docker-v2-manifest-v0.0.1")
	assert.NoError(t, err)
	assert.Contains(t, output, "org.opencontainers.image.version")
	assert.Equal(t, "0.0.1", output["org.opencontainers.image.version"])
}

func TestDockerRetrieveDockerV2ManifestListMetadata(t *testing.T) {
	reg := DockerRegistry{
		Url: "https://ghcr.io",
	}
	output, err := reg.RetrieveMetadata("airfocusio/git-ops-update-test", "docker-v2-manifest-list-v0.0.1")
	assert.NoError(t, err)
	assert.Contains(t, output, "org.opencontainers.image.version")
	assert.Equal(t, "0.0.1", output["org.opencontainers.image.version"])
}

func TestDockerRetrieveOciV1ManifestMetadata(t *testing.T) {
	reg := DockerRegistry{
		Url: "https://ghcr.io",
	}
	output, err := reg.RetrieveMetadata("airfocusio/git-ops-update-test", "oci-v1-manifest-v0.0.1")
	assert.NoError(t, err)
	assert.Contains(t, output, "org.opencontainers.image.version")
	assert.Equal(t, "0.0.1", output["org.opencontainers.image.version"])
}

func TestDockerRetrieveOciV1ImageIndexMetadata(t *testing.T) {
	reg := DockerRegistry{
		Url: "https://ghcr.io",
	}
	output, err := reg.RetrieveMetadata("airfocusio/git-ops-update-test", "oci-v1-image-index-v0.0.1")
	assert.NoError(t, err)
	assert.Contains(t, output, "org.opencontainers.image.version")
	assert.Equal(t, "0.0.1", output["org.opencontainers.image.version"])
}
