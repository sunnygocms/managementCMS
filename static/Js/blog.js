$(document).ready(function() {
	$("#newsform").validate({
		ignore : "",
		rules : {
			title : "required",
			desc	: {
				required:true,
				minlength : 30,
				maxlength:100,
			},
			content : {
				required : true,
				minlength : 30
			}
		},
		messages : {
			title : "请输入文章标题",
			desc	: {
				required : "请输入描述",
				minlength : "描述不能够太短",
				maxlength:"描述不能够太长",
			},
			content : {
				required : "请输入文章正文",
				minlength : "正文不能够太短"
			}
		}
	});

	$('#newsform').submit(function() {
		if ($("#newsform").valid()) {
			return true;
		}

	});
	
	
});


