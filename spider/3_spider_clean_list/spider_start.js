var send = require("./utils/send");
var config = require("./utils/config");
var cheerio = require("cheerio"); //DOM处理插件
var moment = require("moment"); //时间处理插件
var url = require("url"); //URL处理插件

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
		if (data[i].clean_rule_list) {
			eval(data[i].clean_rule_list);
		}
		send.post(config.db_server_url + config.db_list_source_edit, editparam);
	} catch (e) {
		console.error('PARSE_LIST_ERROR[' + data[i]._id + ']:' + data[i].list_url + '\n',e);
		editparam.status = 99;
		send.post(config.db_server_url + config.db_list_source_edit, editparam);
	}
}
//获取绝对链接
function getAbsUrl(baseUrl,href) {
	if(!/^http/i.test(href)){
		if(/^\/\//.test(href)){
			href=url.parse(baseUrl).protocol + href
		}else{ href=url.resolve(baseUrl,href); }
	}
	return href;
}
//发送解析结果函数
function sendAdd(href, title) {
	console.log("{ href:\"%s\", title:\"%s\" }", href, title);
	if (param && title && href) {
		param.detail_url = getAbsUrl(param.list_url,href);
		param.title = title;
		param.timestamp = new Date().getTime();
		send.post(config.db_server_url + config.db_list_detail_source_add, param);
	}else{ console.warn("[WARNING]%s:%s", param.list_url, "列表有空数据"); }
}
