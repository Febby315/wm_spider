const url = require("url"); //URL处理插件
const moment = require("moment"); //时间处理插件
const cheerio = require("cheerio"); //DOM处理插件
const send = require("./utils/send");
const config = require("./utils/config");

//获取urlList表内容的接口
var result = send.post(config.db_server_url + config.db_list_source_query, { conf: { status: 1, web_type: { $ne: "price" } } });
var data = result.data;
var param;
for (let i = 0, len = data.length; i < len; i++) {
	let editparam = { status: 10, _id:data[i]._id, version:data[i].version };
	param = config.clone(data[i]);
	param.dealStatus = 0;
	param.parent_id = data[i]._id;
	try {
		$ = cheerio.load(data[i].list_content, { decodeEntities: false });
		/**
		 * TODO 这里用动态函数替代
		 * new Function("param",data[i].clean_rule_list).bind(param)(param);
		 */
		data[i].clean_rule_list?eval(data[i].clean_rule_list):null;
		send.post(config.db_server_url + config.db_list_source_edit, editparam);
	} catch (err) {
		console.error("[PARSE_LIST_ERROR]:%s (%s)\n%s",data[i].list_url,data[i]._id,err);
		editparam.status = 99;
		send.post(config.db_server_url + config.db_list_source_edit, editparam);
	}
}
//发送解析结果函数
function sendAdd(href, title) {
	console.log(JSON.stringify({ href: href,title: title }));
	if (param && title && href) {
		param.detail_url = url.resolve(param.list_url,href);
		param.title = title;
		param.timestamp = new Date().getTime();
		send.post(config.db_server_url + config.db_list_detail_source_add, param);
	}else{ console.warn("[WARNING]%s:%s", param.list_url, "列表有空数据"); }
}
