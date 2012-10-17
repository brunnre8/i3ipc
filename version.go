package i3ipc

import (
	"encoding/json"
)

type I3Version struct {
	Major          int32
	Minor          int32
	Patch          int32
	Human_Readable string
}

func (self *IPCSocket) GetVersion() (version I3Version, err error) {
	json_reply, err := self.Raw(I3GetVersion, "")
	if err != nil {
		return
	}

	err = json.Unmarshal(json_reply, &version)
	return
}
