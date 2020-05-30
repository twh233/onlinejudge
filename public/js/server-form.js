(function(){
    $("input[type=\"checkbox\"], input[type=\"radio\"]").not("[data-switch-no-init]").bootstrapSwitch();

    // set status
    var val = $('[name="Status"]').val();
    var boxes = $("[data-status-value]");
    for (var i = 0; i < boxes.length; i++) {
        var box = $(boxes[i]);
        box.bootstrapSwitch("state", val&box.data("status-value"));
    }

    $("[data-status-value]").on("switchChange.bootstrapSwitch",function(){
        var val = 0;
        var boxes = $("[data-status-value]");
        for (var i = 0; i < boxes.length; i++) {
            var box = $(boxes[i]);
            if(box.bootstrapSwitch("state")) {
                val = val | box.data("status-value");
            }
        }
        $('[name="Status"]').val(val);
    });

    var Tags = $('[data-stag-value]');

    var get_array = function(args) {
        var tmp = new Map();
        for (var i = 0; i < args.length; i++){
            var t = $(args[i]).data("stag-value");
            tmp.set(t, true)
        }

        return tmp;
    };

    var s = get_array(Tags);
    var in_array = function(tags, tag) {
        for (var i = 0; i < tags.length; i++){
            if (tags[i] == tag) {
                return true;
            }
        }

        return false;
    }

    var tags = $("[data-tags-value]");
    for (var i = 0; i < tags.length; i++) {
        var tag = $(tags[i]);
        var has = s.get(tag.data("tags-value"));
        if (has != true) {
            has = false;
        }
        tag.bootstrapSwitch("state", has)
    }

    var flush_tags_value = function() {
        var tmp = new Array();
        var tags = $("[data-tags-value]");
        for (var i = 0; i < tags.length; i++) {
            var tag = $(tags[i]);
            if (tag.bootstrapSwitch("state")) {
                tmp.push(tag.data("tags-value"))
            }
        }
        
        $('[name="Tags"]').val(tmp);
    };

    flush_tags_value()

    $("[data-tags-value]").on("switchChange.bootstrapSwitch", flush_tags_value)

    var timestamp2date = function(timestamp) {
        return moment.unix(timestamp);
    };

    var date2timestamp = function(d) {
        return d.utc().unix();
    };

    $('#opentime_datetimepicker').datetimepicker({
        format: "YYYY/MM/DD HH:mm:ss",
        sideBySide: true,
    });

    // datepicker 事件
    $('#opentime_datetimepicker').on('dp.change', function(e) {
        var ts = date2timestamp(e.date);
        $('[Name="Opentime"]').val(ts);
    });

    // 使用Opentime初始化datepicker
    var opentime = Number($('[Name="Opentime"]').val());
    if (opentime > 0) {
        var d = timestamp2date(opentime);
        $('#opentime_datetimepicker').data("DateTimePicker").date(d);
    }
})(this);
