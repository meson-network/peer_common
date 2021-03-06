package peer_common

import "github.com/coreservice-io/utils/hash_util"

const MesonAccessKeyMark = "_m_access_key_"

// from Origin_url to Origin_hash
func GenFileHash(pullZone string, requestPath string) string {
	return hash_util.SHA256String(pullZone + requestPath)[0:16]
}
