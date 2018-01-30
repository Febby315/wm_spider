const moment = require('moment');
const cheerio = require('cheerio');
const send = require('./utils/send');
const config = require('./utils/config');

var result = send.post(config.db_server_url + config.db_article_query, { conf:{dealStatus: 0}});
var data = result.data;
for (let i = 0, len = data.length; i < len; i++) {
	var param = config.clone(data[i]);
	param.keyStatus = '0';
	param.nlpDate = moment().format('YYYY-MM-DD HH:mm:ss');
	/**
	 * JSON.stringify([param.articleNavTitle,param.articleClassification]);
	*/
	param.keyWordList = '[' + [param.articleNavTitle,param.articleClassification].join(",") + ']';
	param.keyTitleList = '[' + param.articleNavTitle + ']';
    let editparam = { _id: data[i]._id, version: data[i].version, dealStatus: 10 };
	if(param.info_flag==="sxnyw_trsq"){
		var privateField={};
		var content=param.articleContent.replace(/\s+/gm,"").replace(/（[\d\.]+%-[\d\.]+%）/gm,"");
		var reg=/据.+(渭北).(关中).(陕南).(陕北)?.+（(0-20cm)）.+平均值分别为([\d\.]+)%?.([\d\.]+)%?.([\d\.]+)%?.(([\d\.]+)%?.)?（(20-40cm)）分别为([\d\.]+)%?.([\d\.]+)%?.([\d\.]+)%?.(([\d\.]+)%.)?/gm.exec(content);
		var soilMoisture=privateField.soilMoisture={};
		if(reg){
			soilMoisture.zone=[]; soilMoisture.cm0_20=[]; soilMoisture.cm20_40=[];soilMoisture.unit="%";
			soilMoisture.date=moment(param.articleCreateDateTime).subtract(2,"days").format("YYYY-MM-DD");
			for(var n=1;n<5;n++){
				if(reg[n]){
					if(n===4){
						soilMoisture.zone.push(reg[n]); soilMoisture.cm0_20.push(parseFloat(reg[n+5+1])); soilMoisture.cm20_40.push(parseFloat(reg[n+11+1]));
					}else{
						soilMoisture.zone.push(reg[n]); soilMoisture.cm0_20.push(parseFloat(reg[n+5]));  soilMoisture.cm20_40.push(parseFloat(reg[n+11]));
					}
				}
			}
		}
		param.privateField=JSON.stringify(privateField,null,"\t");
	}
	param.article_id = data[i]._id;
	send.post(config.db_server_url + config.db_tagArticle_add, param);
	send.post(config.db_server_url + config.db_article_edit, editparam);//处理完的数据修改状态
}