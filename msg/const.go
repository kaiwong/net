package msg

const (
	PKG_CRC32_SIZE = 4
)

const (
	PKG_CRC32_BEGIN = 0
	PKG_CRC32_END   = PKG_CRC32_BEGIN + PKG_CRC32_SIZE

	PKG_HEADER_SIZE
)

const (
	MSG_TYPE_SIZE = 1
	UDP_TYPE_SIZE = 1
	MSG_SEQ_SIZE  = 4
	MSG_LEN_SIZE  = 4

	MAX_MESSAGE_SIZE = 10240
)

const (
	MSG_HEADER_BEGIN = 0
	MSG_TYPE_BEGIN
	MSG_TYPE_END = MSG_TYPE_BEGIN + MSG_TYPE_SIZE
	MSG_SEQ_BEGIN
	MSG_SEQ_END = MSG_SEQ_BEGIN + MSG_SEQ_SIZE
	MSG_LEN_BEGIN
	MSG_LEN_END = MSG_LEN_BEGIN + MSG_LEN_SIZE
	MSG_HEADER_END

	MSG_HEADER_SIZE
)

const (
	TYPE_NORMAL = 0x01
	TYPE_FEC    = 0x02
	TYPE_SYN    = 0x03
	TYPE_ACK    = 0x80
	TYPE_PING   = 0x81
	TYPE_PONG   = 0x82
	TYPE_FIN    = 0x83
)

const (
	MSG_STATUS_INIT = 1 << iota
	MSG_STATUS_TRANSMITTED
	MSG_STATUS_ACKED
	MSG_STATUS_LOSS
	MSG_STATUS_CANCEL
)

// ack msg index
const (
	ACK_HEADER_BEGIN = 0
	ACK_TYPE_BEGIN
	ACK_TYPE_END = ACK_TYPE_BEGIN + MSG_TYPE_SIZE
	ACK_SEQ_BEGIN
	ACK_SEQ_END = ACK_SEQ_BEGIN + MSG_SEQ_SIZE
	ACK_NEXT_SEQ_BEGIN
	ACK_NEXT_SEQ_END = ACK_NEXT_SEQ_BEGIN + MSG_SEQ_SIZE
	ACK_ACKED_SEQ_BEGIN
	ACK_ACKED_SEQ_END = ACK_ACKED_SEQ_BEGIN + MSG_SEQ_SIZE
	ACK_HEADER_END

	ACK_HEADER_SIZE
)

const (
	UDP_HEADER_BEGIN = 0
	UDP_TYPE_BEGIN
	UDP_TYPE_END = UDP_TYPE_BEGIN + MSG_TYPE_SIZE
	UDP_SEQ_BEGIN
	UDP_SEQ_END = UDP_SEQ_BEGIN + MSG_SEQ_SIZE
	UDP_LEN_BEGIN
	UDP_LEN_END = UDP_LEN_BEGIN + MSG_LEN_SIZE
	UDP_ACK_SEQ_BEGIN
	UDP_ACK_SEQ_END = UDP_ACK_SEQ_BEGIN + MSG_SEQ_SIZE
	UDP_ACK_NEXT_SEQ_BEGIN
	UDP_ACK_NEXT_SEQ_END = UDP_ACK_NEXT_SEQ_BEGIN + MSG_SEQ_SIZE
	UDP_ACK_ACKED_SEQ_BEGIN
	UDP_ACK_ACKED_SEQ_END = UDP_ACK_ACKED_SEQ_BEGIN + MSG_SEQ_SIZE
	UDP_PADDING           = UDP_ACK_ACKED_SEQ_END + 11
	UDP_HEADER_END

	UDP_HEADER_SIZE
)
