$(document).ready(function() {
	$("#form").validate({
		ignore : "",
		rules : {
			name : "required",
			symbol : "required",
			desc : {
				required : true,
				minlength : 0,
				maxlength : 180
			},
			content : {
				required : true,
				minlength : 0
			}
		},
		messages : {
			name : "请输入频道名称",
			symbol : "请输入频道简称",
			desc : {
				required : "请输入描述",
				minlength : "描述不能够太短",
				maxlength : "描述不能太长"
			},
			content : {
				required : "请输入文章正文",
				minlength : "正文不能够太短"
			}
		}
	});

	$('#form').submit(function() {
		if ($("#form").valid()) {
			return true;
		}

	});
	
	
});


