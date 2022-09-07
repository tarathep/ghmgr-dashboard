package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets7360c1c70ea553c84313cdffe90d9701dfc45025 = "h2{color: blue;}\r\nh1{color: red;}"
var _Assetsf9253fa4eefef29a2a92ddaf9150e0e29d670a5e = "<!doctype html>\r\n<body>\r\n  <p>Hello, {{.Foo}}</p>\r\n  <p><span id=\"foo\"></span> from Javascript</p>\r\n  <script src=\"assets/main.js\"></script>\r\n</body>\r\n"
var _Assetsa9aae58170dcb89c438b811b7c9afd39ff2ba085 = "document.querySelector('#foo').innerHTML = 'Hello!'\r\n"
var _Assets11d2bf3219d6112bf8f2a7bc9eea268d5a74ff54 = "<!DOCTYPE html>\r\n<html lang=\"en\">\r\n \r\n<head>\r\n    <meta charset=\"UTF-8\" />\r\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\" />\r\n    <title>My Table</title>\r\n    <script src=\"https://cdnjs.cloudflare.com/ajax/libs/jquery/3.5.1/jquery.min.js\"></script>\r\n    <script src=\"https://cdn.datatables.net/1.10.21/js/jquery.dataTables.min.js\"></script>\r\n    <link rel=\"stylesheet\" type=\"text/css\" href=\"https://cdn.datatables.net/1.10.21/css/jquery.dataTables.min.css\" />\r\n</head>\r\n \r\n<body>\r\n    <table id=\"myTable\" class=\"display\" style=\"width: 100%;\">\r\n        <thead>\r\n            <tr>\r\n                <th>Name</th>\r\n                <th>Position</th>\r\n                <th>Office</th>\r\n                <th>Age</th>\r\n                <th>Start date</th>\r\n                <th>Salary</th>\r\n            </tr>\r\n        </thead>\r\n        <tbody>\r\n            <tr>\r\n                <td>Tiger Nixon</td>\r\n                <td>System Architect</td>\r\n                <td>Edinburgh</td>\r\n                <td>61</td>\r\n                <td>2011/04/25</td>\r\n                <td>$320,800</td>\r\n            </tr>\r\n            <tr>\r\n                <td>Garrett Winters</td>\r\n                <td>Accountant</td>\r\n                <td>Tokyo</td>\r\n                <td>63</td>\r\n                <td>2011/07/25</td>\r\n                <td>$170,750</td>\r\n            </tr>\r\n            <tr>\r\n                <td>Ashton Cox</td>\r\n                <td>Junior Technical Author</td>\r\n                <td>San Francisco</td>\r\n                <td>66</td>\r\n                <td>2009/01/12</td>\r\n                <td>$86,000</td>\r\n            </tr>\r\n            <tr>\r\n                <td>Cedric Kelly</td>\r\n                <td>Senior Javascript Developer</td>\r\n                <td>Edinburgh</td>\r\n                <td>22</td>\r\n                <td>2012/03/29</td>\r\n                <td>$433,060</td>\r\n            </tr>\r\n            <tr>\r\n                <td>Airi Satou</td>\r\n                <td>Accountant</td>\r\n                <td>Tokyo</td>\r\n                <td>33</td>\r\n                <td>2008/11/28</td>\r\n                <td>$162,700</td>\r\n            </tr>\r\n        </tbody>\r\n    </table>\r\n    <script>\r\n        $(document).ready(function () {\r\n            $(\"#myTable\").DataTable();\r\n        });\r\n    </script>\r\n</body>\r\n \r\n</html>"
var _Assetsa2b1a90d4eee20247b65680898ef00a62c9ef33f = "<!DOCTYPE html>\r\n<html lang=\"en\">\r\n<head>\r\n<meta charset=\"utf-8\">\r\n<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\r\n<title>Bootstrap User Management Data Table</title>\r\n<link rel=\"stylesheet\" href=\"https://fonts.googleapis.com/css?family=Roboto|Varela+Round\">\r\n<!-- <link rel=\"stylesheet\" href=\"https://stackpath.bootstrapcdn.com/bootstrap/4.5.0/css/bootstrap.min.css\"> -->\r\n\r\n<!-- <link rel=\"stylesheet\" href=\"https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/4.5.2/css/bootstrap.css\"> -->\r\n<!-- <link rel=\"stylesheet\" href=\"https://cdn.datatables.net/1.12.1/css/dataTables.bootstrap4.min.css\"> -->\r\n\r\n\r\n<link rel=\"stylesheet\" href=\"https://fonts.googleapis.com/icon?family=Material+Icons\">\r\n<link rel=\"stylesheet\" href=\"https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css\">\r\n\r\n\r\n<script src=\"https://code.jquery.com/jquery-3.5.1.min.js\"></script>\r\n<script src=\"https://cdn.jsdelivr.net/npm/popper.js@1.16.0/dist/umd/popper.min.js\"></script>\r\n<script src=\"https://stackpath.bootstrapcdn.com/bootstrap/4.5.0/js/bootstrap.min.js\"></script>\r\n\r\n\r\n\r\n<link rel=\"stylesheet\" href=\"https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/4.5.2/css/bootstrap.css\">\r\n<link rel=\"stylesheet\" href=\"https://cdn.datatables.net/1.12.1/css/dataTables.bootstrap4.min.css\">\r\n\r\n\r\n\r\n<script src=\"https://code.jquery.com/jquery-3.5.1.js\"></script>\r\n<script src=\"https://cdn.datatables.net/1.12.1/js/jquery.dataTables.min.js\"></script>\r\n<script src=\"https://cdn.datatables.net/1.12.1/js/dataTables.bootstrap4.min.js\"></script>\r\n\r\n\r\n<style>\r\nbody {\r\n    color: #566787;\r\n    background: #f5f5f5;\r\n    font-family: 'Varela Round', sans-serif;\r\n    font-size: 13px;\r\n}\r\n.table-responsive {\r\n    margin: 10px 0;\r\n}\r\n.table-wrapper {\r\n    min-width: 1500px;\r\n    background: #fff;\r\n    padding: 20px 25px;\r\n    border-radius: 3px;\r\n    box-shadow: 0 1px 1px rgba(0,0,0,.05);\r\n}\r\n.table-title {\r\n    padding-bottom: 15px;\r\n    background: #299be4;\r\n    color: #fff;\r\n    padding: 16px 30px;\r\n    margin: -20px -25px 10px;\r\n    border-radius: 3px 3px 0 0;\r\n}\r\n.table-title h2 {\r\n    margin: 5px 0 0;\r\n    font-size: 24px;\r\n}\r\n.table-title .btn {\r\n    color: #566787;\r\n    float: right;\r\n    font-size: 13px;\r\n    background: #fff;\r\n    border: none;\r\n    min-width: 50px;\r\n    border-radius: 2px;\r\n    border: none;\r\n    outline: none !important;\r\n    margin-left: 10px;\r\n}\r\n.table-title .btn:hover, .table-title .btn:focus {\r\n    color: #566787;\r\n    background: #f2f2f2;\r\n}\r\n.table-title .btn i {\r\n    float: left;\r\n    font-size: 21px;\r\n    margin-right: 5px;\r\n}\r\n.table-title .btn span {\r\n    float: left;\r\n    margin-top: 2px;\r\n}\r\ntable.table tr th, table.table tr td {\r\n    border-color: #e9e9e9;\r\n    padding: 12px 15px;\r\n    vertical-align: middle;\r\n}\r\ntable.table tr th:first-child {\r\n    width: 60px;\r\n}\r\ntable.table tr th:last-child {\r\n    width: 100px;\r\n}\r\ntable.table-striped tbody tr:nth-of-type(odd) {\r\n    background-color: #fcfcfc;\r\n}\r\ntable.table-striped.table-hover tbody tr:hover {\r\n    background: #f5f5f5;\r\n}\r\ntable.table th i {\r\n    font-size: 13px;\r\n    margin: 0 5px;\r\n    cursor: pointer;\r\n}\t\r\ntable.table td:last-child i {\r\n    opacity: 0.9;\r\n    font-size: 22px;\r\n    margin: 0 5px;\r\n}\r\ntable.table td a {\r\n    font-weight: bold;\r\n    color: #566787;\r\n    display: inline-block;\r\n    text-decoration: none;\r\n}\r\ntable.table td a:hover {\r\n    color: #2196F3;\r\n}\r\ntable.table td a.settings {\r\n    color: #2196F3;\r\n}\r\ntable.table td a.mode_edit {\r\n    color: #f3bb21;\r\n}\r\ntable.table td a.delete {\r\n    color: #F44336;\r\n}\r\ntable.table td i {\r\n    font-size: 19px;\r\n}\r\ntable.table .avatar {\r\n    border-radius: 50%;\r\n    vertical-align: middle;\r\n    margin-right: 10px;\r\n}\r\n.status {\r\n    font-size: 30px;\r\n    margin: 2px 2px 0 0;\r\n    display: inline-block;\r\n    vertical-align: middle;\r\n    line-height: 10px;\r\n}\r\n.text-success {\r\n    color: #10c469;\r\n}\r\n.text-info {\r\n    color: #62c9e8;\r\n}\r\n.text-warning {\r\n    color: #FFC107;\r\n}\r\n.text-danger {\r\n    color: #ff5b5b;\r\n}\r\n.pagination {\r\n    float: right;\r\n    margin: 0 0 5px;\r\n}\r\n.pagination li a {\r\n    border: none;\r\n    font-size: 13px;\r\n    min-width: 30px;\r\n    min-height: 30px;\r\n    color: #999;\r\n    margin: 0 2px;\r\n    line-height: 30px;\r\n    border-radius: 2px !important;\r\n    text-align: center;\r\n    padding: 0 6px;\r\n}\r\n.pagination li a:hover {\r\n    color: #666;\r\n}\t\r\n.pagination li.active a, .pagination li.active a.page-link {\r\n    background: #03A9F4;\r\n}\r\n.pagination li.active a:hover {        \r\n    background: #0397d6;\r\n}\r\n.pagination li.disabled i {\r\n    color: #ccc;\r\n}\r\n.pagination li i {\r\n    font-size: 16px;\r\n    padding-top: 6px\r\n}\r\n.hint-text {\r\n    float: left;\r\n    margin-top: 10px;\r\n    font-size: 13px;\r\n}\r\n</style>\r\n\r\n\r\n</head>\r\n<body>\r\n<div class=\"container-fluid\">\r\n    <div class=\"table-responsive\">\r\n        <div class=\"table-wrapper\">\r\n            <div class=\"table-title\">\r\n                <div class=\"row\">\r\n                    <div class=\"col-sm-5\">\r\n                        <h2>GHMGR <b>User Management</b></h2>\r\n                    </div>\r\n                    <div class=\"col-sm-7\">\r\n                        <a href=\"#\" class=\"btn btn-secondary\"><i class=\"material-icons\">&#xE147;</i> <span>Add New User</span></a>\r\n                        <a href=\"#\" class=\"btn btn-secondary\"><i class=\"material-icons\">&#xE24D;</i> <span>Export to Excel</span></a>\t\t\t\t\t\t\r\n                    </div>\r\n                </div>\r\n            </div>\r\n\r\n            <table id=\"myTable\" class=\"table table-striped table-bordered\" style=\"width:100%\">\r\n                <thead>\r\n                    <tr>\r\n                        <th>#</th>\r\n                        <th>Username</th>\r\n                        <th>Full-Name</th>\t\t\t\t\t\t\r\n                        <th>Email</th>\r\n                        <th>Role</th>\r\n                        <th>Subscription Owner</th>\r\n                        <th>GitHub Username</th>\r\n                        <th>GitHub Role</th>\r\n                        <th>GitHub Status</th>\r\n                        <th>Azure DEV Status</th>\r\n                        <th>Azure PRD Status</th>\r\n                        <th>ELK Status</th>\r\n                        <th>Jumphost Status</th>\r\n                        <th>Bastion Status</th>\r\n                    </tr>\r\n                </thead>\r\n                <tbody>\r\n                      <tr>\r\n                        <td>#</td>\r\n                        <td><a href=\"#\"><img src=\"https://avatars.githubusercontent.com/u/26413446?v=4\" width=\"50%\" class=\"avatar\" alt=\"Avatar\"> tarathec</a></td>\r\n                        <td>Tarathep Choengchaiyaphum</td>\t\t\t\t\t\t\r\n                        <td>tarathec@ais.co.th</td>\r\n                        <td>cdc application modernize/devsecops</td>\r\n                        <td>Dev sub owner</td>\r\n                        <td>tarathec</td>\r\n                        <td>maintainer</td>\r\n                        <td><span class=\"status text-warning\">&bull;</span> Active,Not Validate email</td>\r\n                        <td><span class=\"status text-success\">&bull;</span> Active</td>\r\n                        <td><span class=\"status text-success\">&bull;</span> Active</td>\r\n                        <td><span class=\"status text-success\">&bull;</span> Active</td>\r\n                        <td><span class=\"status text-success\">&bull;</span> Active</td>\r\n                        <td><span class=\"status text-danger\">&bull;</span> Active</td>\r\n                    </tr>\r\n                    <!-- <tr>\r\n                        <td>1</td>\r\n                        <td><a href=\"#\"><img src=\"https://avatars.githubusercontent.com/u/26413446?v=4\" width=\"10%\" class=\"avatar\" alt=\"Avatar\"> Tarathep Choengchaiyaphum</a></td>\r\n                        <td>04/10/2013</td>                        \r\n                        <td>Admin</td>\r\n                        <td><span class=\"status text-success\">&bull;</span> Active</td>\r\n                        <td>\r\n                            <a href=\"#\" class=\"mode_edit\" title=\"Settings\" data-toggle=\"tooltip\"><i class=\"material-icons\">&#xF097;</i></a>\r\n                            <a href=\"#\" class=\"delete\" title=\"Delete\" data-toggle=\"tooltip\"><i class=\"material-icons\">&#xE5C9;</i></a>\r\n                        </td>\r\n                    </tr> -->\r\n                    <!-- <tr>\r\n                        <td>2</td>\r\n                        <td><a href=\"#\"><img src=\"/examples/images/avatar/2.jpg\" class=\"avatar\" alt=\"Avatar\"> Paula Wilson</a></td>\r\n                        <td>05/08/2014</td>                       \r\n                        <td>Publisher</td>\r\n                        <td><span class=\"status text-success\">&bull;</span> Active</td>\r\n                        <td>\r\n                            <a href=\"#\" class=\"settings\" title=\"Settings\" data-toggle=\"tooltip\"><i class=\"material-icons\">&#xE8B8;</i></a>\r\n                            <a href=\"#\" class=\"delete\" title=\"Delete\" data-toggle=\"tooltip\"><i class=\"material-icons\">&#xE5C9;</i></a>\r\n                        </td>\r\n                    </tr>\r\n                    <tr>\r\n                        <td>3</td>\r\n                        <td><a href=\"#\"><img src=\"/examples/images/avatar/3.jpg\" class=\"avatar\" alt=\"Avatar\"> Antonio Moreno</a></td>\r\n                        <td>11/05/2015</td>\r\n                        <td>Publisher</td>\r\n                        <td><span class=\"status text-danger\">&bull;</span> Suspended</td>                        \r\n                        <td>\r\n                            <a href=\"#\" class=\"settings\" title=\"Settings\" data-toggle=\"tooltip\"><i class=\"material-icons\">&#xE8B8;</i></a>\r\n                            <a href=\"#\" class=\"delete\" title=\"Delete\" data-toggle=\"tooltip\"><i class=\"material-icons\">&#xE5C9;</i></a>\r\n                        </td>                        \r\n                    </tr>\r\n                    <tr>\r\n                        <td>4</td>\r\n                        <td><a href=\"#\"><img src=\"/examples/images/avatar/4.jpg\" class=\"avatar\" alt=\"Avatar\"> Mary Saveley</a></td>\r\n                        <td>06/09/2016</td>\r\n                        <td>Reviewer</td>\r\n                        <td><span class=\"status text-success\">&bull;</span> Active</td>\r\n                        <td>\r\n                            <a href=\"#\" class=\"settings\" title=\"Settings\" data-toggle=\"tooltip\"><i class=\"material-icons\">&#xE8B8;</i></a>\r\n                            <a href=\"#\" class=\"delete\" title=\"Delete\" data-toggle=\"tooltip\"><i class=\"material-icons\">&#xE5C9;</i></a>\r\n                        </td>\r\n                    </tr>\r\n                    <tr>\r\n                        <td>5</td>\r\n                        <td><a href=\"#\"><img src=\"https://avatars.githubusercontent.com/u/26413447?v=4\" class=\"avatar\" alt=\"Avatar\" width=\"25px\"> Bokie Tarathep</a></td>\r\n                        <td>12/08/2017</td>                        \r\n                        <td>Moderator</td>\r\n                        <td><span class=\"status text-warning\">&bull;</span> Email not validate</td>\r\n                        <td>\r\n                            <a href=\"#\" class=\"settings\" title=\"Settings\" data-toggle=\"tooltip\"><i class=\"material-icons\">&#xE8B8;</i></a>\r\n                            <a href=\"#\" class=\"delete\" title=\"Delete\" data-toggle=\"tooltip\"><i class=\"material-icons\">&#xE5C9;</i></a>\r\n                        </td>\r\n                    </tr>\r\n                    <tr>\r\n                      <td>6</td>\r\n                      <td><a href=\"#\"><img src=\"https://avatars.githubusercontent.com/u/26413446?v=4\" class=\"avatar\" alt=\"Avatar\" width=\"25px\"> Tarathep Choengchaiyaphum</a></td>\r\n                      <td>12/08/2017</td>                        \r\n                      <td>Moderator</td>\r\n                      <td><span class=\"status text-warning\">&bull;</span> Email not validate</td>\r\n                      <td>\r\n                          <a href=\"#\" class=\"settings\" title=\"Settings\" data-toggle=\"tooltip\"><i class=\"material-icons\">&#xE8B8;</i></a>\r\n                          <a href=\"#\" class=\"delete\" title=\"Delete\" data-toggle=\"tooltip\"><i class=\"material-icons\">&#xE5C9;</i></a>\r\n                      </td>\r\n                  </tr> -->\r\n                </tbody>\r\n            </table>\r\n\r\n            <script>\r\n              $(document).ready(function () {\r\n                  $(\"#myTable\").DataTable();\r\n              });\r\n            </script>\r\n\r\n        </div>\r\n    </div>\r\n</div>\r\n\r\n\r\n<div class=\"modal fade\" id=\"orangeModalSubscription\" tabindex=\"-1\" role=\"dialog\" aria-labelledby=\"myModalLabel\"\r\n  aria-hidden=\"true\">\r\n  <div class=\"modal-dialog modal-notify modal-warning\" role=\"document\">\r\n    <!--Content-->\r\n    <div class=\"modal-content\">\r\n      <!--Header-->\r\n      <div class=\"modal-header text-center\">\r\n        <h4 class=\"modal-title white-text w-100 font-weight-bold py-2\">Subscribe</h4>\r\n        <button type=\"button\" class=\"close\" data-dismiss=\"modal\" aria-label=\"Close\">\r\n          <span aria-hidden=\"true\" class=\"white-text\">&times;</span>\r\n        </button>\r\n      </div>\r\n\r\n      <!--Body-->\r\n      <div class=\"modal-body\">\r\n        <div class=\"md-form mb-5\">\r\n          <i class=\"fas fa-user prefix grey-text\"></i>\r\n          <input type=\"text\" id=\"form3\" class=\"form-control validate\">\r\n          <label data-error=\"wrong\" data-success=\"right\" for=\"form3\">Your name</label>\r\n        </div>\r\n\r\n        <div class=\"md-form\">\r\n          <i class=\"fas fa-envelope prefix grey-text\"></i>\r\n          <input type=\"email\" id=\"form2\" class=\"form-control validate\">\r\n          <label data-error=\"wrong\" data-success=\"right\" for=\"form2\">Your email</label>\r\n        </div>\r\n      </div>\r\n\r\n      <!--Footer-->\r\n      <div class=\"modal-footer justify-content-center\">\r\n        <a type=\"button\" class=\"btn btn-outline-warning waves-effect\">Send <i class=\"fas fa-paper-plane-o ml-1\"></i></a>\r\n      </div>\r\n    </div>\r\n    <!--/.Content-->\r\n  </div>\r\n</div>\r\n\r\n<div class=\"text-center\">\r\n  <a href=\"\" class=\"btn btn-default btn-rounded\" data-toggle=\"modal\" data-target=\"#orangeModalSubscription\">Launch\r\n    modal Subscription</a>\r\n</div>\r\n\r\n</body>\r\n</html>"
var _Assetsc5cc2dadf26cb5b4a53904786db07d3095062cf7 = "<!doctype html>\r\n<head>\r\n<link href=\"/assets/c.css\" rel=\"stylesheet\">\r\n</head>\r\n<body>\r\n  <p>Can you see this? → {{.Bar}}</p>\r\n  <h1>COLOR</h1>\r\n</body>\r\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"assets"}, "/assets": []string{"bar.tmpl", "c.css", "index.tmpl", "main.js", "table.tmpl", "template.tmpl"}}, map[string]*assets.File{
	"/assets": &assets.File{
		Path:     "/assets",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1662540333, 1662540333068053700),
		Data:     nil,
	}, "/assets/bar.tmpl": &assets.File{
		Path:     "/assets/bar.tmpl",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1658144955, 1658144955067758200),
		Data:     []byte(_Assetsc5cc2dadf26cb5b4a53904786db07d3095062cf7),
	}, "/assets/c.css": &assets.File{
		Path:     "/assets/c.css",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1658145352, 1658145352557616400),
		Data:     []byte(_Assets7360c1c70ea553c84313cdffe90d9701dfc45025),
	}, "/assets/index.tmpl": &assets.File{
		Path:     "/assets/index.tmpl",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1658142929, 1658142929886644700),
		Data:     []byte(_Assetsf9253fa4eefef29a2a92ddaf9150e0e29d670a5e),
	}, "/assets/main.js": &assets.File{
		Path:     "/assets/main.js",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1658142929, 1658142929886644700),
		Data:     []byte(_Assetsa9aae58170dcb89c438b811b7c9afd39ff2ba085),
	}, "/assets/table.tmpl": &assets.File{
		Path:     "/assets/table.tmpl",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1658392361, 1658392361521291700),
		Data:     []byte(_Assets11d2bf3219d6112bf8f2a7bc9eea268d5a74ff54),
	}, "/assets/template.tmpl": &assets.File{
		Path:     "/assets/template.tmpl",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1662539939, 1662539939872810800),
		Data:     []byte(_Assetsa2b1a90d4eee20247b65680898ef00a62c9ef33f),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1662540751, 1662540751863492600),
		Data:     nil,
	}}, "")
