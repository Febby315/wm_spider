
var request = require('sync-request');
var send = {
		post : function(url,form) {
			var res = request('POST', url, {
			  json: form
			});
			return  JSON.parse(res.getBody('utf8'));
		},
		get: function(url,form) {
		var res = request('GET', url, {
		  'headers': {
			'user-agent': 'example-user-agent'
		  }
		});
		return  res.getBody();
		}
}
module.exports = send;

//---------------------------------------------
