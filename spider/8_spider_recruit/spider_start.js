var send = require('./utils/send');
var config = require('./utils/config');
//查找配置文件
var result = send.post(config.db_server_url + config.db_list_detail_cleaned_query, { conf:{ 'dealStatus': 2, "web_type": "recruit" } });
var data = result.data;
for (var i in data) {
	let param = { _id: data[i]._id, version: data[i].version, dealStatus: 10 };
	try {
		let addparam = config.clone(data[i].table_info);
		addparam.title = data[i].title;
		addparam.content=data[i].content;
		addparam.publishTime = data[i].pub_time;
		addparam.contentUrl = data[i].detail_url;
		addparam.timestamp = data[i].timestamp;
		addparam.pickTime = data[i].operate_date;
		send.post(config.db_server_url + config.db_recruit_add, addparam);
		send.post(config.db_server_url + config.db_list_detail_cleaned_edit, param);
	} catch (e) {
		console.error("新增失败", e);
		param.dealStatus = 99;
		send.post(config.db_server_url + config.db_list_detail_cleaned_edit, param);
	}
}