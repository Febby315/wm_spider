const url = require("url");
const moment = require("moment");
const cheerio = require("cheerio");
const send = require("./utils/send");
const config = require("./utils/config");
// 获取需要处理的任务
var result = send.post(config.db_server_url + config.db_list_detail_source_query, { conf: { dealStatus: 1 } });
var data = result.data;
var param;
for (let i = 0, len = data.length; i < len; i++) {
	let editparam = { dealStatus: 10, _id:data[i]._id, version:data[i].version };
	param = config.clone(data[i]);
	param.dealStatus = 0;
	param.parent_id = data[i]._id;
	if (data[i] && data[i].content_sour && data[i].clean_rule_content) {	//对象,对象.内容,对象.清洗规则
		try {	// 执行清洗
			/**
			 * TODO 这里用动态函数替代
			 * new Function("param",data[i].clean_rule_list).bind(param)(param);
			 */
			$ = cheerio.load(data[i].content_sour, { decodeEntities: false });
			data[i].clean_rule_content?eval(data[i].clean_rule_content):null;
			sendAdd();
			send.post(config.db_server_url + config.db_list_detail_source_edit, editparam);
		} catch (err) {	// 异常处理并通知服务器
			console.error("[PARSE_DETAIL_ERROR]:%s (%s)\n%s",data[i].detail_url,data[i]._id,err);
			editparam.dealStatus = 99;
			send.post(config.db_server_url + config.db_list_detail_source_edit, editparam);
		}
	}
}
// 解析结果返回服务器
function sendAdd() {
	if (param && param.content_sour) {
		send.post(config.db_server_url + config.db_list_detail_cleaned_add, param);
	}else{ console.warn("[WARNING]%s:%s", param.detail_url, "详情有空数据"); }
}