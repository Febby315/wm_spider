var cheerio = require("cheerio"); //jquery 方式操作
var send = require("./utils/send");
var config = require("./utils/config");
//查找配置文件
var result = send.post(config.db_server_url + config.db_list_detail_cleaned_query,{ conf: { dealStatus: 1 } });
var data = result.data;
for (let i = 0, len = data.length; i < len; i++) {
	let param = { _id: data[i]._id, version: data[i].version, dealStatus: 2 };
	try {
		var contentImageSour = JSON.parse(data[i].contentImageSour || "[]");
		var pageContent = data[i].content_sour;
		if (pageContent) {
			$ = cheerio.load(pageContent, { decodeEntities: false });
			$("img").each(function(ei, element) {
				var imgsrc = $(element).attr("src");
				if (!imgsrc) { imgsrc = $(element).attr("data-src"); }
				for (var j=0;j<contentImageSour.length;j++) {
					if (contentImageSour[j].sour_url.indexOf(imgsrc.replace(/(^\/|\.\/|\.\.\/)/gm,""))!==-1) {
						$(element).attr("src", contentImageSour[j].url).attr("onerror", "this.style.display='none';");
					}else{ $(element).attr("onerror", "this.style.display='none';"); }
				}
			});
			param.content = $.html();
		}
		//新增记录
		send.post(config.db_server_url + config.db_list_detail_cleaned_edit, param);
	} catch (e) {
		console.error(e);
		param.dealStatus = 99;
		send.post(config.db_server_url + config.db_list_detail_cleaned_edit, param);
	}
}
