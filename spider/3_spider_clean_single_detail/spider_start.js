var send = require("./utils/send");
var config = require("./utils/config");
var cheerio = require("cheerio"); //DOM处理插件
var moment = require("moment"); //时间处理插件
var url = require("url"); //URL处理插件

//增加一不,先将table_info数据塞满,再将数据放到price表中
var result = send.post(config.db_server_url + config.db_list_source_query, { conf: { status: 1, web_type: "price" } });
var data = result.data;
var param;
for (let i = 0, len = data.length; i < len; i++) {	//遍历配置表获取过来的数据
	let editparam = { status: 10, _id:data[i]._id, version:data[i].version };
	param = config.clone(data[i].table_info);
	param.dealStatus = 0;
	param.parent_id = data[i]._id;
	try {
		$ = cheerio.load(data[i].list_content, { decodeEntities: false });
		if (data[i].clean_rule_content) {
			eval(data[i].clean_rule_content);
		}
		send.post(config.db_server_url + config.db_list_source_edit, editparam);
	} catch (e) {
		console.error('PARSE_PRICE_ERROR[' + data[i]._id + ']:' + data[i].list_url + '\n',e);
		editparam.status = 99;
		send.post(config.db_server_url + config.db_list_source_edit, editparam);
	}
}
function sendAdd() {	//发布时间转换成时间戳
	try {
		param.timestamp = new Date(param.publishDate).getTime();
	} catch (e) {
		param.timestamp = new Date().getTime();
	}
	send.post(config.db_server_url + config.db_price_add, param);
}