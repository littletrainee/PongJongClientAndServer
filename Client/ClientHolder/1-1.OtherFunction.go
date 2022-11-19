package ClientHolder

import Enum "github.com/littletrainee/PongJong_Client_And_Server/Client/EnumHolder"

func (ch *ClientHolder) GetURL(index Enum.Pattern) string {
	return "http://" + ch.Address.Get() + ":" + ch.Port.Get() + "/" + ch.Pattern[index]
}
