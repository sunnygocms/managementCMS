$(document).ready(function() {
	$("#newsform").validate({
		ignore : "",
		rules : {
			recommend_id : {
					digits:true,
					required : true,
					maxlength :6
				},
			recommend_name : "required",
			symbol: "required"

		},
		messages : {
			recommend_id : {
				digits: "只能输入数字",
				required : "请输分类ID",
				maxlength :"分类ID长度为6"
			},
			recommend_name : "请输入标题",
			symbol : "请输入英文简称"
		}
	});
	
	$('#newsform').submit(function() {
		if ($("#newsform").valid()) {
			return true;
		}
	});
});	