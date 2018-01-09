var send = require("./utils/send");
var config = require("./utils/config");
var cheerio = require("cheerio");	//DOM处理插件
var moment = require("moment");	//时间处理插件
var url = require("url");	//URL处理插件
//查找配置文件
var result = send.post(config.db_server_url + config.db_list_detail_source_query, { conf: { dealStatus: 1 } });
var data = result.data;
var param;
for (let i = 0, len = data.length; i < len; i++) {
	let editparam = { dealStatus: 10, _id:data[i]._id, version:data[i].version };
	param = config.clone(data[i]);
	param.dealStatus = 0;
	param.parent_id = data[i]._id;
	if (data[i] && data[i].content_sour && data[i].clean_rule_content) {	//对象,对象.内容,对象.清洗规则
		try {
			$ = cheerio.load(data[i].content_sour, { decodeEntities: false });
			eval(data[i].clean_rule_content);
			sendAdd();
			send.post(config.db_server_url + config.db_list_detail_source_edit, editparam);
		} catch (e) {
			console.error("PARSE_DETAIL_ERROR[" + data[i]._id + "]:" + data[i].detail_url + "\n", e);
			editparam.dealStatus=99;
			send.post(config.db_server_url + config.db_list_detail_source_edit, editparam);
		}
	}
}
function sendAdd() {
	if (param && param.content_sour) {
		send.post(config.db_server_url + config.db_list_detail_cleand_add, param);
	}else{ console.warn("[WARNING]%s:%s", param.detail_url, "详情有空数据"); }
}