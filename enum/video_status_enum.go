package enum

import "annie-baby/utils"

var VideoStatusEnumName = "VideoStatus"

var VideoEnum = utils.NewEnum(VideoStatusEnumName)

var Pending, _ = VideoEnum.Set(1, "PENDING")
var Info, _ = VideoEnum.Set(2, "INFO")
var Crawling, _ = VideoEnum.Set(3, "CRAWLING")
var CrawFail, _ = VideoEnum.Set(4, "CRAWFAIL")
var CrawSucc, _ = VideoEnum.Set(5, "CRAWSUCC")
