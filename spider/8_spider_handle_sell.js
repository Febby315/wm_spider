const send = require("./utils/send");
const config = require("./utils/config");

var result = send.post(config.db_server_url + config.db_list_detail_cleaned_query,{ conf: { dealStatus: 2, web_type: "sell" } });
var data = result.data;
for (let i = 0, len = data.length; i < len; i++) {
	let param = { _id: data[i]._id, version: data[i].version, dealStatus: 10 };
	try {
		let addparam = config.clone(data[i].table_info);
		addparam.parent_id = data[i]._id;
		addparam.contentUrl = data[i].detail_url;
		try {
			addparam.timestamp = new Date(data[i].pub_time).getTime();
		} catch (e) {
			addparam.timestamp = new Date().getTime();
		}
		addparam.publishTime = data[i].pub_time;
		addparam.title = data[i].title;
		addparam.content = data[i].content;
		addparam.pickTime = data[i].operate_date;
		send.post(config.db_server_url + config.db_sell_add, addparam);
		send.post(config.db_server_url + config.db_list_detail_cleaned_edit, param);
	} catch (err) {
		console.error("[HANDLE_sell_ERROR]:%s\n%s",data[i]._id,err);
		param.dealStatus = 99;
		send.post(config.db_server_url + config.db_list_detail_cleaned_edit, param);
	}
	
}
