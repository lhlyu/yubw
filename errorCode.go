package errorCode

import (
	"errorCode/i18n"
)

const Version = "1.0.0"

const (
	ZH_CN = iota
	ES_UN
)

var CodeMap map[int]string

func init(){
	SetLG(ZH_CN)
}

func SetLG(lg int){
	switch lg {
	case 0:
		CodeMap = i18n.ZHCNCodeMap
	case 1:
		CodeMap = i18n.ESUNCodeMap
	default:
		CodeMap = i18n.ZHCNCodeMap
	}
}



