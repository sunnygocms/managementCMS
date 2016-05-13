$(function() {
    var action_log = {add:[],edit:[],remove:[]};
    var editor = $("<div>").tooltip({
        position : "bottom center",
        offset : [2, 0],
        events : {
            def : "xxx,ooo",
            tooltip : "xxx,ooo"
        },
        tip : "#f_r_editor"
    });
    var loading = $("<a>").overlay({
        // some mask tweaks suitable for modal dialogs
        mask : {
            color : '#fff',
            loadSpeed : 200,
            opacity : 0.8
        },
        closeOnClick : false,
        target : $("#loading")
    }).overlay();
    var $dialog = $("#f_r_editor");
    var dialog = $("<a>").overlay({
        // some mask tweaks suitable for modal dialogs
        mask : {
            color : '#fff',
            loadSpeed : 200,
            opacity : 0.8
        },
        closeOnClick : false,
        onBeforeLoad : function(event) {
            setData();
        },
        onLoad : function(event) {
            if(dialog.id){
                $dialog.find("#f_r_editor_short_title").focus();
            }
            else{
                $dialog.find(".long_title").focus();
            }
            dialog.target.addClass("current_edit");
        },
        onClose:function(){
            dialog.target.removeClass("current_edit");
        },
        target : $dialog
    }).overlay();
    function setData(data) {
        if (!data) {
            data = dialog;
        }
        $dialog.find(".id").text(data.type + ":" + data.id)
        $dialog.find(".long_title").val(data.long_title);
        $dialog.find("#f_r_editor_short_title").val(data.short_title);
    };
    //重置
    $dialog.find(".reset").click(function() {
        setData();
    });
    //保存
    $dialog.find(".ok").click(function() {
        var shortTitle = $dialog.find("#f_r_editor_short_title").val();
        if(shortTitle.length){
            var newsId = $dialog.find(".id").text();
            var recommendId = dialog.target.parents(".f_r:first").attr("id").split("_")[2];
            var $a,isAdd = dialog.target.is(".add");
            if(isAdd){
                $a = $("<a>");
            }
            else{
                $a = dialog.target;
            }
            $a.html($dialog.find("#f_r_editor_short_title").val())
                .attr("title",$dialog.find(".long_title").val())
                .attr("id","f_rn_"+newsId.split(":")[1]+"_"+recommendId)
                .attr("data-type",newsId.split(":")[0]);
            if(isAdd){
                var $li = $("<"+dialog.target.prop("nodeName")+" class='f_rn_li'>");
                $li.append($a.addClass("f_rn"));
                $li.append("<span class='remove'>x</span>");
                $li.insertAfter(dialog.target);
                //这里先把原来的克隆，然后插在原有位置，再把原始的删掉，再将新的启用sortable
                //原因是新插入的li没有绑定事件，无法拖动
                //如果直接重新sortable有可能重复绑定事件
                //clone出的新东西是没有绑定任何事件的，所以可以对新的li做sortable
                var $parent = $li.parents(".f_r");
                $parent.after($parent=$parent.clone()).remove();
                $parent.sortable({
                    items:"li:not(.add),dd:not(.add)",
                    forcePlaceholderSize:true
                });
            }
            dialog.close();
        }
        else{
            alert("请填写标题");
            $dialog.find("#f_r_editor_short_title").focus();
        }
    });
    //autocomplete
    $dialog.find(".long_title").keyup(function() {
        var val = $(this).val();
        if (val.length > 1) {
            var oldVal = $(this).data("oldVal");
            if (val != oldVal) {
                $.getJSON(site_url + "/admin.php?s=/Future/searchNews", {
                    key : val
                }).done(function(result) {
                    var list = result.data;
                    if (list) {
                        var $ul = $dialog.find(".autocomplete").empty().show();
                        for (var i in list) {
                            $("<li>").text(list[i].long_title).data("data", list[i]).appendTo($ul);
                        }
                    }
                });
                $(this).data("oldVal", val);
            }
        }
    });
    $dialog.on("click",function(event){
        var $target = $(event.target); 
        if($target.is($dialog.find(".long_title"))
        ||$target.is($dialog.find(".autocomplete"))
        ||$dialog.find(".autocomplete").has($target).length){
            
        }
        else{
            $dialog.find(".autocomplete").hide();
        }
        //   
    })
    //显示/隐藏边界
    $("#borderToggleBtn").toggle(function(){
        $(".f_r").addClass("border-toggle");
    },function(){
        $(".f_r").removeClass("border-toggle");
    });
    //保存
    $("#saveBtn").click(function(){
        var data = [];
        $(".f_r").each(function(){
            var recommendId = $(this).attr("id").substr(4);
            var recommendData = {
                recommendId:$(this).attr("id").substr(4),
                list:[]
            };
            var dt = $(this).find("dt").text();
            if(dt){
                recommendData.recommendName = dt;
            }
            $(this).find(".f_rn_li a").each(function(){
                var newsData = {
                    id:$(this).attr("id").split("_")[2],
                    short_title:$(this).html(),
                    type:$(this).attr("data-type")
                };
                recommendData.list.push(newsData);
            });
            data.push(recommendData);
            
        });
        loading.load();
        $.post(site_url+"/admin.php?s=/Future/homepage&submit=true",{
            data:JSON.stringify(data)
        },function(result){
            if(result.status){
                alert("保存成功");
            }
            else{
                alert("保存失败：\n"+result.info);
            }
        },"json").always(function(){
            loading.close();
        }).fail(function(){
            alert("网络异常");
        });
    });
    //拖动排序控件
    $(".f_r").sortable({
        items:"li:not(.add),dd:not(.add)",
        forcePlaceholderSize:true
    });
    $(document).on("sortupdate",".f_r",function(){
    });
    //搜索文章的下拉框
    $(document).on("click", ".autocomplete li:not(.empty)", function() {
        var data = $(this).data("data");
        $dialog.find(".autocomplete").hide();
        if(data.short_title==null){
            data.short_title = data.long_title;
        }
        setData(data);
    });
    //文章标题的a标签
    $(document).on("click", ".f_rn", function() {
        dialog.type = $(this).attr("data-type");
        dialog.id = $(this).attr("id").split("_")[2];
        dialog.long_title = $(this).attr("title");
        dialog.short_title = $(this).html();
        dialog.target = $(this);
        dialog.load();
    });
    //添加新文章按钮
    $(document).on("click", ".f_r .add", function() {
        dialog.type = null;
        dialog.id = null;
        dialog.long_title = null;
        dialog.short_title = null;
        dialog.target = $(this);
        dialog.load();
    });
    //删除推荐的x按钮
    $(document).on("click", ".f_rn_li .remove", function() {
        if (confirm("确认删除？")) {
            var id = $(this).prev(".f_rn").attr("id").substr(5);
            action_log.remove.push(id);
            $(this).parent().remove();
        }
    });
    $(document).on("click",".f_r dt",function(){
        var val = prompt("请输入新的标题",$(this).html());
        if(val){
            val = $.trim(val);
            if(val.length){
                $(this).html(val);
            }
        }
    });
}); 