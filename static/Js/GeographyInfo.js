$(document).ready(function() {
	$("#form").validate({
		ignore : "",
		rules : {
			name : "required",
			provinceID : "required",
			
		},
		messages : {
			name : "请输入区县名称名称",
			provinceID : "属于哪个省不能为空"
		}
	});

	$('#form').submit(function() {
		if ($("#form").valid()) {
			return true;
		}

	});
	
	
});


