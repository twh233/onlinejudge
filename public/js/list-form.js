    function update_page(fservers, page) {
        if (page <= 1) {
            $("#up").attr('disabled', true)
        } else {
            $("#up").attr('disabled', false)
        }

        var tag = $("#tag").val();
        if (page >= get_page_max(fservers, tag)){
            $("#down").attr('disabled', true)
        } else {
            $("#down").attr('disabled', false)
        }

        $("#page").html(page);

        //show(tag, page) no to do html script
    };


    page_num = 30
    function get_page_max(fservers, tag){
        var servers = fservers[tag];
        var tmp = servers.length / page_num;

        if (tmp != parseInt(tmp)) {
            tmp = parseInt(tmp) + 1
        }
        return tmp
    };


    function get_servers(fservers, tag, page){
        var servers = fservers[tag];
        var tmp = new Array();
        for (var i = (page-1)*page_num; i < page*page_num;i++){
            if (i >= servers.length) {
                continue;
            };

            tmp.push(servers[i])
        };

        return tmp
    };

    function fmt_servers(servers, tags){
        var fservers = {};
        fservers["all"] = servers;
        for (k in tags) {
            fservers[k] = new Array();
        };

        for (k in servers) {
            s = servers[k];
            for (i in s.Tags) {
                fservers[s.Tags[i]].push(s)
            }
        };

        return fservers
    };
