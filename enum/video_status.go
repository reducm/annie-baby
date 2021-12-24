package enum

type VideoEnum int64

const (
	PENDING       VideoEnum = 1
	INFO          VideoEnum = 2
	CRAWL_FAIL    VideoEnum = 3
	CRAWL_SUCCESS VideoEnum = 4
)

func (s VideoEnum) String() string {
	switch s {
	case PENDING:
		return "PENDING"
	case INFO:
		return "INFO"
	case CRAWL_FAIL:
		return "CRAWL_FAIL"
	case CRAWL_SUCCESS:
		return "CRAWL_SUCCESS"
	}
	return "UNKNOW"
}
