var send = require("./utils/send");
var config = require("./utils/config");
var moment = require("moment");
var operdate = moment().format("YYYY-MM-DD HH:mm:ss");
let ishave = true;
let currentPage = 1;
while (ishave) {
	//确保数据取完
	var result = send.post(config.db_server_url + config.db_spider_config_query, { currentPage, conf: { status: 0 } });
	var data = result.data;
	currentPage++;
	if (data.length == 0) {
		ishave = false;
	}
	for (var i in data) {
		var more_page_model = data[i].more_page_model;
		if (data[i].pick_ways && data[i].pick_ways == "sogou") {
			var from_url = data[i].from_url; //
			var url = from_url.split(/[ |,;]/);
			var query_conf = data[i].query_conf;
			var conf = query_conf.split(/[ |,;]/);
			// console.log(from_url);
			for (let from_url_num in url) {
				//console.log(from_url_num,'fromurl ',from_url[from_url_num].url_path,from_url[from_url_num].url_name);
				for (var num in conf) {
					var param = config.clone(data[i]);
					param.current_page = data[i].page_count;
					param.status = 0;
					param.operate_date = operdate;
					param.from_url = url[from_url_num];
					param.query_conf = conf[num];
					param.parent_id = data[i]._id;
					send.post(
						config.db_server_url + config.db_spider_list_source_add,
						param
					);
				}
			}
		} else if (more_page_model && more_page_model.indexOf("{PAGE}") > -1) {
			for (var num = 1; num <= data[i].page_count; num++) {
				var param = config.clone(data[i]);
				param.current_page = num;
				param.status = 0;
				param.operate_date = operdate;
				param.list_url = more_page_model.replace("{PAGE}", num);
				param.parent_id = data[i]._id;
				send.post(
					config.db_server_url + config.db_spider_list_source_add,
					param
				);
			}
		} else {
			var param = config.clone(data[i]);
			param.current_page = data[i].page_count;
			param.status = 0;
			param.operate_date = operdate;
			param.list_url = more_page_model;
			param.parent_id = data[i]._id;
			send.post(config.db_server_url + config.db_spider_list_source_add, param);
		}
	}
}