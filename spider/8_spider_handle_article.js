const send = require("./utils/send");
const config = require("./utils/config");

var result = send.post(config.db_server_url + config.db_list_detail_cleaned_query,{ conf: { dealStatus: 2, web_type: "article" } });
var data = result.data;
for (let i = 0, len = data.length; i < len; i++) {
	let param = { _id: data[i]._id, version: data[i].version, dealStatus: 10 };
	try {
		let addparam = config.clone(data[i].table_info);
		addparam.parent_id = data[i]._id;
		addparam.articleURL = data[i].detail_url;
		addparam.articleCreateDateTime = data[i].pub_time;
		addparam.articleAbstract = data[i].summary;
		addparam.articleTitle = data[i].title;
		addparam.articleContent = data[i].content;
		addparam.nlpDate = data[i].operate_date;
		addparam.articleRefineTime = data[i].operate_date;
		addparam.articleRefineTimestamp = data[i].timestamp; //爬取时间戳
		addparam.articleImageList = data[i].listShowImage;
		addparam.ext_column_value = JSON.stringify(addparam.ext_column_value||"{}");
		send.post(config.db_server_url + config.db_article_add, addparam);
		send.post(config.db_server_url + config.db_list_detail_cleaned_edit, param);
	} catch (err) {
		console.error("[HANDLE_article_ERROR]:%s\n%s",data[i]._id,err);
		param.dealStatus = 99;
		send.post(config.db_server_url + config.db_list_detail_cleaned_edit, param);
	}
}
