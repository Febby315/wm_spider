const moment = require("moment");
const send = require("./utils/send");
const config = require("./utils/config");

let operdate = moment().format("YYYY-MM-DD HH:mm:ss");
let ishave = true,currentPage = 1;
while (ishave) {
	//确保数据取完
	var result = send.post(config.db_server_url + config.db_spider_config_query, { currentPage, conf: { status: 0 } });
	var data = result.data;
	ishave = data.length > 0;
	for (var i in data) {
		var more_page_model = data[i].more_page_model;
		if (data[i].pick_ways && data[i].pick_ways == "sogou") {	// 判断是否为搜狗数据源
			var url = data[i].from_url.split(/[ |,;]/);
			var conf = data[i].query_conf.split(/[ |,;]/);
			for (let from_url_num in url) {
				for (var num in conf) {
					var param = config.clone(data[i]);
					param.current_page = data[i].page_count;
					param.status = 0;
					param.operate_date = operdate;
					param.from_url = url[from_url_num];
					param.query_conf = conf[num];
					param.parent_id = data[i]._id;
					send.post(config.db_server_url + config.db_list_source_add, param);
				}
			}
		} else if (more_page_model && more_page_model.indexOf("{PAGE}") > -1) {	// 判断是否为分页模板
			for (var num = 1; num <= data[i].page_count; num++) {
				var param = config.clone(data[i]);
				param.current_page = num;
				param.status = 0;
				param.operate_date = operdate;
				param.list_url = more_page_model.replace("{PAGE}", num);
				param.parent_id = data[i]._id;
				send.post(config.db_server_url + config.db_list_source_add, param);
			}
		} else {	// 非分页模板
			var param = config.clone(data[i]);
			param.current_page = data[i].page_count;
			param.status = 0;
			param.operate_date = operdate;
			param.list_url = more_page_model;
			param.parent_id = data[i]._id;
			send.post(config.db_server_url + config.db_list_source_add, param);
		}
	}
	// 当前分页++
	currentPage++;
}