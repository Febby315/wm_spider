const url = require("url");
const moment = require("moment");
const cheerio = require("cheerio");
const send = require("./utils/send");
const config = require("./utils/config");
// 获取需要处理的任务
var result = send.post(config.db_server_url + config.db_list_source_query, { conf: { status: 1, web_type: { $ne: "price" } } });
var data = result.data;
var param;
for (let i = 0, len = data.length; i < len; i++) {
	let editparam = { status: 10, _id:data[i]._id, version:data[i].version };
	param = config.clone(data[i]);
	param.dealStatus = 0;
	param.parent_id = data[i]._id;
	try {	// 执行清洗
		/**
		 * TODO 这里用动态函数替代
		 * new Function("param",data[i].clean_rule_list).bind(param)(param);
		 */
		$ = cheerio.load(data[i].list_content, { decodeEntities: false });
		data[i].clean_rule_list?eval(data[i].clean_rule_list):null;
		send.post(config.db_server_url + config.db_list_source_edit, editparam);
	} catch (err) {	// 异常处理并通知服务器
		console.error("[PARSE_LIST_ERROR]:%s (%s)\n%s",data[i].list_url,data[i]._id,err);
		editparam.status = 99;
		send.post(config.db_server_url + config.db_list_source_edit, editparam);
	}
}
// 解析结果返回服务器
function sendAdd(href, title) {
	console.log(JSON.stringify({ href: href,title: title }));
	if (param && title && href) {
		param.detail_url = url.resolve(param.list_url,href);
		param.title = title;
		param.timestamp = new Date().getTime();
		send.post(config.db_server_url + config.db_list_detail_source_add, param);
	}else{ console.warn("[WARNING]%s:%s", param.list_url, "列表有空数据"); }
}
