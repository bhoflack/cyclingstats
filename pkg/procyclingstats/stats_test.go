package procyclingstats

import (
	"testing"
)

const content = `
<!DOCTYPE HTML>
<html>
<head>
    <!-- Global site tag (gtag.js) - Google Analytics -->
    <script async src="https://www.googletagmanager.com/gtag/js?id=G-2MNJ20S7SJ"></script>
    <script>
    window.dataLayer = window.dataLayer || [];
    function gtag() {
        dataLayer.push(arguments);
    }
    gtag('js', new Date());

    gtag('config', 'G-2MNJ20S7SJ');
    </script>

    <title>Wout van Aert </title>
    <base href="https://www.procyclingstats.com/rider.php"/>
    <link rel="canonical" href="https://www.procyclingstats.com/rider/wout-van-aert"/>
    <script>
    var r89 = r89 || {
        callAds: [],
        pushAd: function(id, name, config) {
            r89.callAds.push([id, name, config]);
        }
    };
    </script>
    <script async src="https://tags.refinery89.com/v2/procyclingstats.js"></script>
    <meta name="description" content="Wout van Aert (born 1994-09-15 in Herentals) is a professional road racing cyclist from Belgium, currently riding for Team Visma | Lease a Bike. His best results are 9 stage wins in Tour de France and 2 wins in E3 Saxo Classic.">
    <meta name="keywords" content="">
    <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
    <meta charset="utf-8">
    <meta name="facebook-domain-verification" content="tjwn3jyjupm0t1j8e0sgis3c80x88r"/>

    <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.4.1/jquery.min.js"></script>
    <script src="https://ajax.googleapis.com/ajax/libs/jqueryui/1.12.1/jquery-ui.min.js"></script>
    <base href="rider.php"/>
    <link rel="stylesheet" type="text/css" href="https://www.procyclingstats.com/pcs_a202.css"/>
    <link rel="stylesheet" type="text/css" href="https://www.procyclingstats.com/pcs_b168.css"/>
    <script src="https://www.procyclingstats.com/pcs_s82.js"></script>
    <link rel="stylesheet" href="//code.jquery.com/ui/1.12.0/themes/base/jquery-ui.css">
    <link rel="apple-touch-icon" sizes="180x180" href="https://www.procyclingstats.com/images/pcs-logos/icon3/apple-touch-icon.png">
    <link rel="icon" type="image/png" sizes="32x32" href="https://www.procyclingstats.com/images/pcs-logos/icon3/favicon-32x32.png">
    <link rel="icon" type="image/png" sizes="16x16" href="https://www.procyclingstats.com/images/pcs-logos/icon3/favicon-16x16.png">
    <link rel="shortcut icon" type="image/x-icon" href="https://www.procyclingstats.com/images/pcs-logos/icon3/favicon.ico"/>

</head>
<body>
    <div class="wrapper">
        <header class="">
            <div class="cont">
                <a class="logo" href="index.php"></a>

                <form action="search.php" class="msearch">
                    <input type="hidden" id="term"/>
                    <input type="text" name="term" id="search" data-nav="rider/-id-/start" data-nav2="" data-page="rider" class="search4" placeholder="search"/>
                </form>
                <div class="menuCont">
                    <a class="mobMenu" href="">Menu&#x25BC;</a>
                    <ul class="nav_v4 default">
                        <li class=" ">
                            <a class="reg" href="index.php">Home</a>
                        </li>
                        <li class="more4 ">
                            <a class="reg" href="races.php">Races</a>
                            <a class="more4 toggleNavDropdown" href="">&#x25BC;</a>
                            <ul class="hide">
                                <li>
                                    <a href="race/tour-de-france">Tour de France</a>
                                </li>
                                <li>
                                    <a href="race/giro-d-italia">Giro d'Italia</a>
                                </li>
                                <li>
                                    <a href="race/vuelta-a-espana">La Vuelta ciclista a España</a>
                                </li>
                                <li>
                                    <a href="race/world-championship">World Championships</a>
                                </li>
                                <li>
                                    <a href="race/amstel-gold-race">Amstel Gold Race</a>
                                </li>
                                <li>
                                    <a href="race/milano-sanremo">Milano-Sanremo</a>
                                </li>
                                <li>
                                    <a href="race/tirreno-adriatico">Tirreno-Adriatico</a>
                                </li>
                                <li>
                                    <a href="race/liege-bastogne-liege">Liège-Bastogne-Liège</a>
                                </li>
                                <li>
                                    <a href="race/il-lombardia">Il Lombardia</a>
                                </li>
                                <li>
                                    <a href="race/la-fleche-wallone">La Flèche Wallonne</a>
                                </li>
                                <li>
                                    <a href="race/paris-nice">Paris - Nice</a>
                                </li>
                                <li>
                                    <a href="race/paris-roubaix">Paris-Roubaix</a>
                                </li>
                                <li>
                                    <a href="race/volta-a-catalunya">Volta Ciclista a Catalunya</a>
                                </li>
                                <li>
                                    <a href="race/dauphine">Critérium du Dauphiné</a>
                                </li>
                                <li>
                                    <a href="race/ronde-van-vlaanderen">Tour des Flandres</a>
                                </li>
                                <li>
                                    <a href="race/gent-wevelgem">Gent-Wevelgem in Flanders Fields</a>
                                </li>
                                <li>
                                    <a href="race/san-sebastian">Clásica Ciclista San Sebastián</a>
                                </li>
                            </ul>
                        </li>
                        <li class="more4 ">
                            <a class="reg" href="teams.php">Teams</a>
                            <a class="more4 toggleNavDropdown" href="">&#x25BC;</a>
                            <ul class="hide">
                                <li>
                                    <a href="team/ineos-grenadiers-2024">INEOS Grenadiers</a>
                                </li>
                                <li>
                                    <a href="team/groupama-fdj-2024">Groupama - FDJ</a>
                                </li>
                                <li>
                                    <a href="team/ef-education-easypost-2024">EF Education-EasyPost</a>
                                </li>
                                <li>
                                    <a href="team/decathlon-ag2r-la-mondiale-2024">Decathlon AG2R La Mondiale Team</a>
                                </li>
                                <li>
                                    <a href="team/cofidis-2024">Cofidis</a>
                                </li>
                                <li>
                                    <a href="team/bora-hansgrohe-2024">BORA - hansgrohe</a>
                                </li>
                                <li>
                                    <a href="team/bahrain-victorious-2024">Bahrain - Victorious</a>
                                </li>
                                <li>
                                    <a href="team/astana-qazaqstan-team-2024">Astana Qazaqstan Team</a>
                                </li>
                                <li>
                                    <a href="team/intermarche-circus-want-2024">Intermarché - Wanty</a>
                                </li>
                                <li>
                                    <a href="team/lidl-trek-2024">Lidl - Trek</a>
                                </li>
                                <li>
                                    <a href="team/movistar-team-2024">Movistar Team</a>
                                </li>
                                <li>
                                    <a href="team/soudal-quick-step-2024">Soudal - Quick Step</a>
                                </li>
                                <li>
                                    <a href="team/team-dsm-firmenich-postnl-2024">Team dsm-firmenich PostNL</a>
                                </li>
                                <li>
                                    <a href="team/team-jayco-alula-2024">Team Jayco AlUla</a>
                                </li>
                                <li>
                                    <a href="team/team-visma-lease-a-bike-2024">Team Visma | Lease a Bike</a>
                                </li>
                                <li>
                                    <a href="team/uae-team-emirates-2024">UAE Team Emirates</a>
                                </li>
                                <li>
                                    <a href="team/arkea-b-b-hotels-2024">Arkéa - B&B Hotels</a>
                                </li>
                                <li>
                                    <a href="team/alpecin-deceuninck-2024">Alpecin-Deceuninck</a>
                                </li>
                            </ul>
                        </li>
                        <li class=" ">
                            <a class="reg" href="rankings.php">Rankings</a>
                        </li>
                        <li class="more4 ">
                            <a class="reg" href="statistics.php">Statistics</a>
                            <a class="more4 toggleNavDropdown" href="">&#x25BC;</a>
                            <ul class="hide">
                                <li>
                                    <a href="statistics">Overview</a>
                                </li>
                                <li>
                                    <a href="statistics/riders">Riders</a>
                                </li>
                                <li>
                                    <a href="statistics/teams">Teams</a>
                                </li>
                                <li>
                                    <a href="statistics/grandtours">Grand tours</a>
                                </li>
                            </ul>
                        </li>
                        <li class=" ">
                            <a class="reg" href="game.php">Game</a>
                        </li>
                        <li class="more4 ">
                            <a class="reg" href="riders.php">More</a>
                            <a class="more4 toggleNavDropdown" href="">&#x25BC;</a>
                            <ul class="hide">
                                <li>
                                    <a href="riders">Riders</a>
                                </li>
                                <li>
                                    <a href="info/contact">Contact us</a>
                                </li>
                                <li>
                                    <a href="info/countdown-to-1-billion-pageviews">Countdown to 2 billion pageviews</a>
                                </li>
                                <li>
                                    <a href="favorite500.php">Favorite500</a>
                                </li>
                                <li>
                                    <a href="info/sitemap">Sitemap</a>
                                </li>
                                <li>
                                    <a href="info/about-us">About us</a>
                                </li>
                                <li>
                                    <a href="info/profile-score-explained">Profile Score</a>
                                </li>
                                <li>
                                    <a href="info/tutorials">Tutorials</a>
                                </li>
                            </ul>
                        </li>
                        <li class="more4 pro">
                            <a class="reg" href="premium.php">Pro</a>
                            <a class="more4 toggleNavDropdown" href="">&#x25BC;</a>
                            <ul class="hide">
                                <li>
                                    <a href="">Dashboard</a>
                                </li>
                                <li>
                                    <a href="">Lists</a>
                                </li>
                            </ul>
                        </li>

                    </ul>
                </div>
                <a class="user4" data-logged="0" href="account.php">
                    <span>&#10070;</span>
                    <font class="hideIfMobile">login</font>
                </a>
                <a class="settings4 " href="settings.php">
                    <span>&#10033;</span>
                    <font>settings</font>
                </a>
            </div>
        </header>
        <div class="content   ">
            <div class="page-topnav ">
                <ul class="pagenavLinks">
                    <li class="reg ">
                        <a class="cur" href="rider/wout-van-aert">Overview</a>
                    </li>
                    <li class="reg ">
                        <a class="" href="rider/wout-van-aert/statistics">Statistics&#x25BC;</a>
                        <ul>
                            <li>
                                <a href="rider/wout-van-aert/statistics/overview">Overview</a>
                            </li>
                            <li>
                                <a href="rider/wout-van-aert/statistics/season-statistics">Statistics by season</a>
                            </li>
                            <li>
                                <a href="rider/wout-van-aert/statistics/key-events">Key events</a>
                            </li>
                            <li>
                                <a href="rider/wout-van-aert/statistics/world-map-of-rider-races">World map of rider races</a>
                            </li>
                            <li>
                                <a href="rider/wout-van-aert/statistics/wins">Wins</a>
                            </li>
                            <li>
                                <a href="rider/wout-van-aert/statistics/most-starts-by-race">Most starts by race</a>
                            </li>
                            <li>
                                <a href="rider/wout-van-aert/statistics/grand-tour-starts">Starts and results</a>
                            </li>
                            <li>
                                <a href="rider/wout-van-aert/statistics/injury-history">Injury history</a>
                            </li>
                        </ul>
                    </li>
                    <li class="reg ">
                        <a class="" href="rider/wout-van-aert/results">Results</a>
                    </li>
                    <li class="reg ">
                        <a class="" href="rider/wout-van-aert/head-to-head">H2H</a>
                    </li>
                    <li class="reg ">
                        <a class="" href="rider/wout-van-aert/calendar">Calendar</a>
                    </li>
                    <li class="reg ">
                        <a class="" href="rider/wout-van-aert/contribute">More&#x25BC;</a>
                        <ul>
                            <li>
                                <a href="rider/wout-van-aert/contribute/contribute-info">Contribute info</a>
                            </li>
                            <li>
                                <a href="rider/wout-van-aert/contribute/submit-program">Submit program</a>
                            </li>
                            <li>
                                <a href="rider/wout-van-aert/contribute/report">Report</a>
                            </li>
                            <li>
                                <a href="rider/wout-van-aert/contribute/contribute-social-feeds">Contribute sites / social media</a>
                            </li>
                            <li>
                                <a href="rider/wout-van-aert/contribute/contribute-family">Family</a>
                            </li>
                            <li>
                                <a href="rider/wout-van-aert/contribute/embed-rider-results">Embed rider results</a>
                            </li>
                        </ul>
                    </li>
                    <li class="reg ">
                        <a class="" href="rider/wout-van-aert/specialties">Specialties&#x25BC;</a>
                        <ul>
                            <li>
                                <a href="rider/wout-van-aert/specialties/top-gc-results">Top GC results</a>
                            </li>
                            <li>
                                <a href="rider/wout-van-aert/specialties/fastest-tt">Fastest TT</a>
                            </li>
                        </ul>
                    </li>
                    <li class="more">
                        <a href="rider/wout-van-aert">More</a>
                        <ul>
                            <li data-n="1">
                                <a href="rider/wout-van-aert/overview/overview">Overview - Overview</a>
                            </li>
                            <li data-n="2">
                                <a href="rider/wout-van-aert/statistics/statistics">Statistics - Statistics</a>
                            </li>
                            <li data-n="1">
                                <a href="rider/wout-van-aert/statistics/overview">Overview</a>
                            </li>
                            <li data-n="1">
                                <a href="rider/wout-van-aert/statistics/season-statistics">Statistics by season</a>
                            </li>
                            <li data-n="1">
                                <a href="rider/wout-van-aert/statistics/key-events">Key events</a>
                            </li>
                            <li data-n="1">
                                <a href="rider/wout-van-aert/statistics/world-map-of-rider-races">World map of rider races</a>
                            </li>
                            <li data-n="1">
                                <a href="rider/wout-van-aert/statistics/wins">Wins</a>
                            </li>
                            <li data-n="1">
                                <a href="rider/wout-van-aert/statistics/most-starts-by-race">Most starts by race</a>
                            </li>
                            <li data-n="1">
                                <a href="rider/wout-van-aert/statistics/grand-tour-starts">Starts and results</a>
                            </li>
                            <li data-n="1">
                                <a href="rider/wout-van-aert/statistics/injury-history">Injury history</a>
                            </li>
                            <li data-n="3">
                                <a href="rider/wout-van-aert/results/results">Results - Results</a>
                            </li>
                            <li data-n="4">
                                <a href="rider/wout-van-aert/head-to-head/head-to-head">H2H - H2H</a>
                            </li>
                            <li data-n="5">
                                <a href="rider/wout-van-aert/calendar/calendar">Calendar - Calendar</a>
                            </li>
                            <li data-n="6">
                                <a href="rider/wout-van-aert/contribute/contribute">More - More</a>
                            </li>
                            <li data-n="5">
                                <a href="rider/wout-van-aert/contribute/contribute-info">Contribute info</a>
                            </li>
                            <li data-n="5">
                                <a href="rider/wout-van-aert/contribute/submit-program">Submit program</a>
                            </li>
                            <li data-n="5">
                                <a href="rider/wout-van-aert/contribute/report">Report</a>
                            </li>
                            <li data-n="5">
                                <a href="rider/wout-van-aert/contribute/contribute-social-feeds">Contribute sites / social media</a>
                            </li>
                            <li data-n="5">
                                <a href="rider/wout-van-aert/contribute/contribute-family">Family</a>
                            </li>
                            <li data-n="5">
                                <a href="rider/wout-van-aert/contribute/embed-rider-results">Embed rider results</a>
                            </li>
                            <li data-n="7">
                                <a href="rider/wout-van-aert/specialties/specialties">Specialties - Specialties</a>
                            </li>
                            <li data-n="6">
                                <a href="rider/wout-van-aert/specialties/top-gc-results">Top GC results</a>
                            </li>
                            <li data-n="6">
                                <a href="rider/wout-van-aert/specialties/fastest-tt">Fastest TT</a>
                            </li>
                        </ul>
                    </li>
                </ul>
                <div class="clear"></div>
            </div>
            <div class="page-title ">
                <div class="main">
                    <span class="flag be w32"></span>
                    <span class="flag be w24"></span>
                    <h1 class="">Wout van Aert</h1>
                    <span class="hideIfMobile center">&raquo;</span>
                    <span class="red hideIfMobile">Team Visma | Lease a Bike</span>
                </div>
                <div class="sub">
                    <span class="red showIfMobile">Team Visma | Lease a Bike</span>
                </div>
            </div>
            <div class="clear"></div>
            <div class="desktop-ad leaderbord rider"></div>
            <div class="mob-ad leaderbord rider"></div>
            <div class="top-cont"></div>
            <div class="top-cont-mobile top-cont-mobile-rider"></div>
            <div class="page-content page-object default " data-no_menu="1">
                <h2></h2>
                <div class=" ">
                    <ul class="list horizontal fs14 circle bluelink">
                        <li>
                            <a class="" href="rider-in-race/wout-van-aert/omloop-het-nieuwsblad">Rider statistics in Omloop Het Nieuwsblad ME</a>
                        </li>
                    </ul>
                    <div class="left w75 mb_w100">
                        <div class="left w50 mb_w100">
                            <h3>Rider</h3>
                            <div class="rdr-img-cont">
                                <a href="rider/wout-van-aert/statistics/overview">
                                    <img alt="Profile photo of Wout van Aert" src="images/riders/bp/ea/wout-van-aert-2024.jpeg"/>
                                </a>
                            </div>
                            <div class="rdr-info-cont">
                                <b>Date of birth:</b>
                                 15
                                <sup>th</sup>
                                 September 1994 (29)
                                <br/>
                                <b>Nationality:</b>
                                <span class="flag  be"></span>
                                <a class="black" href="nation/belgium">Belgium</a>
                                <br/>
                                <span>
                                    <b>Weight:</b>
                                     78 kg &nbsp; 
                                    <span>
                                        <b>Height:</b>
                                         1.90 m
                                        <br/>
                                        <b>Place of birth:</b>
                                        <a href="location/herentals">Herentals</a>
                                        <br/>
                                        <div class="pps">
                                            <h4>Points per specialty</h4>
                                            <ul>
                                                <li>
                                                    <div class="bar ">
                                                        <div class="bg green" style="width: 100%; height: 100%;  "></div>
                                                    </div>
                                                    <div class="pnt">6832</div>
                                                    <div class="title">
                                                        <a href="rider/wout-van-aert/results/career-points-one-day-races">One day races</a>
                                                    </div>
                                                </li>
                                                <li>
                                                    <div class="bar ">
                                                        <div class="bg red" style="width: 16.774004683841%; height: 100%;  "></div>
                                                    </div>
                                                    <div class="pnt">1146</div>
                                                    <div class="title">
                                                        <a href="rider/wout-van-aert/results/career-points-gc">GC</a>
                                                    </div>
                                                </li>
                                                <li>
                                                    <div class="bar ">
                                                        <div class="bg blue" style="width: 38.173302107728%; height: 100%;  "></div>
                                                    </div>
                                                    <div class="pnt">2608</div>
                                                    <div class="title">
                                                        <a href="rider/wout-van-aert/results/career-points-time-trial">Time trial</a>
                                                    </div>
                                                </li>
                                                <li>
                                                    <div class="bar ">
                                                        <div class="bg orange" style="width: 27.371194379391%; height: 100%;  "></div>
                                                    </div>
                                                    <div class="pnt">1870</div>
                                                    <div class="title">
                                                        <a href="rider/wout-van-aert/results/career-points-sprint">Sprint</a>
                                                    </div>
                                                </li>
                                                <li>
                                                    <div class="bar ">
                                                        <div class="bg purple" style="width: 61.007025761124%; height: 100%;  "></div>
                                                    </div>
                                                    <div class="pnt">4168</div>
                                                    <div class="title">
                                                        <a href="rider/wout-van-aert/results/career-points-climbers">Climber</a>
                                                    </div>
                                                </li>
                                            </ul>
                                        </div>
                                        <div class=""></div>
                                        <div class="mt10">
                                            <ul class="list horizontal sites">
                                                <li>
                                                    <span class="icon twitter left"></span>
                                                    <a target="_blank" class="" href="https://twitter.com/WoutvanAert">twitter</a>
                                                </li>
                                                <li>
                                                    <span class="icon strava left"></span>
                                                    <a target="_blank" class="" href="https://www.strava.com/pros/189040">strava</a>
                                                </li>
                                                <li>
                                                    <span class="icon instagram left"></span>
                                                    <a target="_blank" class="" href="https://www.instagram.com/woutvanaert/">instagram</a>
                                                </li>
                                                <li>
                                                    <span class="icon website left"></span>
                                                    <a target="_blank" class="" href="https://www.woutvanaert.be/">site</a>
                                                </li>
                                                <li>
                                                    <span class="icon facebook left"></span>
                                                    <a target="_blank" class="" href="https://www.facebook.com/woutvaert/">facebook</a>
                                                </li>
                                                <li>
                                                    <span class="icon Cyclo-cross left"></span>
                                                    <a target="_self" class="" href="https://cx.procyclingstats.com/rider/wout-van-aert">Cyclo-cross</a>
                                                </li>
                                            </ul>
                                        </div>
                                        <div class="">
                                            <span class="table-cont">
                                                <ul class="list horizontal rdr-rankings" style=" ">
                                                    <li class="">
                                                        <div class="title   ">
                                                            <a href="rankings/me/all-time">All time</a>
                                                        </div>
                                                        <div class="rnk   ">81</div>
                                                        <div class="clear"></div>
                                                    </li>
                                                    <li class="">
                                                        <div class="title   ">
                                                            <a href="rankings/me/uci-individual">UCI World</a>
                                                        </div>
                                                        <div class="rnk   ">6</div>
                                                        <div class="clear"></div>
                                                    </li>
                                                    <li class="">
                                                        <div class="title   ">
                                                            <a href="rankings/me/individual">PCS Ranking</a>
                                                        </div>
                                                        <div class="rnk   ">6</div>
                                                        <div class="clear"></div>
                                                    </li>
                                                </ul>
                                            </span>
                                        </div>
                                        <div class="">
                                            <span class="table-cont">
                                                <ul class="list horizontal" style=" ">
                                                    <li class="">
                                                        <div class="bold   ">Visits:</div>
                                                        <div>
                                                            <span class="delta-up">&#x25B2;10667</span>
                                                        </div>
                                                        <div>  this week</div>
                                                        <div class="clear"></div>
                                                    </li>
                                                </ul>
                                            </span>
                                        </div>
                            </div>
                            <div class="clear"></div>
                        </div>
                        <div class="right w47 mb_w100">
                            <h3>Top results</h3>
                            <ul class="list moblist flex" data-shortnr="8" style="line-height: 17px; ">
                                <li class="main">
                                    <div class="ar" style="width: 55px; ">
                                        <b>9x</b>
                                        &nbsp;
                                    </div>
                                    <div>
                                        <span class="blue">stage</span>
                                        <a href="race/tour-de-france/2022/stage-20">Tour de France</a>
                                        <span style="color: #777; font: 11px tahoma;">&nbsp; ('22, '21, '20, '19)</span>
                                    </div>
                                </li>
                                <li class="main">
                                    <div class="ar" style="width: 55px; ">
                                        <b>2x</b>
                                        &nbsp;
                                    </div>
                                    <div>
                                        <span class="blue"></span>
                                        <a href="race/e3-harelbeke/2023/result">E3 Saxo Classic</a>
                                        <span style="color: #777; font: 11px tahoma;">&nbsp; ('23, '22)</span>
                                    </div>
                                </li>
                                <li class="main">
                                    <div class="ar" style="width: 55px; ">&nbsp; </div>
                                    <div>
                                        <span class="blue"></span>
                                        <a href="race/milano-sanremo/2020/result">Milano-Sanremo</a>
                                        <span style="color: #777; font: 11px tahoma;">&nbsp; ('20)</span>
                                    </div>
                                </li>
                                <li class="main">
                                    <div class="ar" style="width: 55px; ">&nbsp; </div>
                                    <div>
                                        <span class="blue"></span>
                                        <a href="race/amstel-gold-race/2021/result">Amstel Gold Race</a>
                                        <span style="color: #777; font: 11px tahoma;">&nbsp; ('21)</span>
                                    </div>
                                </li>
                                <li class="main">
                                    <div class="ar" style="width: 55px; ">&nbsp; </div>
                                    <div>
                                        <span class="blue"></span>
                                        <a href="race/strade-bianche/2020/result">Strade Bianche</a>
                                        <span style="color: #777; font: 11px tahoma;">&nbsp; ('20)</span>
                                    </div>
                                </li>
                                <li class="main">
                                    <div class="ar" style="width: 55px; ">&nbsp; </div>
                                    <div>
                                        <span class="blue"></span>
                                        <a href="race/gent-wevelgem/2021/result">Gent-Wevelgem</a>
                                        <span style="color: #777; font: 11px tahoma;">&nbsp; ('21)</span>
                                    </div>
                                </li>
                                <li class="main">
                                    <div class="ar" style="width: 55px; ">&nbsp; </div>
                                    <div>
                                        <span class="blue"></span>
                                        <a href="race/bretagne-classic/2022/result">Bretagne Classic - Ouest-France</a>
                                        <span style="color: #777; font: 11px tahoma;">&nbsp; ('22)</span>
                                    </div>
                                </li>
                                <li class="main">
                                    <div class="ar" style="width: 55px; ">&nbsp; </div>
                                    <div>
                                        <span class="blue"></span>
                                        <a href="race/omloop-het-nieuwsblad/2022/result">Omloop Het Nieuwsblad ME</a>
                                        <span style="color: #777; font: 11px tahoma;">&nbsp; ('22)</span>
                                    </div>
                                </li>
                                <li class="main">
                                    <div class="ar" style="width: 55px; ">
                                        <b>2x</b>
                                        &nbsp;
                                        <span class="shirt st4 w14"></span>
                                    </div>
                                    <div>
                                        <span class="blue">GC</span>
                                        <a href="race/tour-of-britain/2023/gc">Tour of Britain</a>
                                        <span style="color: #777; font: 11px tahoma;">&nbsp; ('23, '21)</span>
                                    </div>
                                </li>
                                <li class="main">
                                    <div class="ar" style="width: 55px; ">
                                        <b>5x</b>
                                        &nbsp;
                                    </div>
                                    <div>
                                        <span class="blue">stage</span>
                                        <a href="race/dauphine/2022/stage-5">Critérium du Dauphiné</a>
                                        <span style="color: #777; font: 11px tahoma;">&nbsp; ('22, '20, '19)</span>
                                    </div>
                                </li>
                                <li class="main">
                                    <div class="ar" style="width: 55px; ">
                                        &nbsp; 
                                        <span class="shirt st5 w14"></span>
                                    </div>
                                    <div>
                                        <span class="blue">Points GC</span>
                                        <a href="race/tour-de-france/2022/stage-21-points">Tour de France</a>
                                        <span style="color: #777; font: 11px tahoma;">&nbsp; ('22)</span>
                                    </div>
                                </li>
                                <li class="main">
                                    <div class="ar" style="width: 55px; ">
                                        <b>2x</b>
                                        &nbsp;2nd&nbsp;
                                    </div>
                                    <div>
                                        <span class="blue"></span>
                                        <a href="race/world-championship/2023/result">World Championships ME - Road Race</a>
                                        <span style="color: #777; font: 11px tahoma;">&nbsp; ('23, '20)</span>
                                    </div>
                                </li>
                            </ul>
                            <h3 class="mt10">Top CX results</h3>
                            <ul class="list moblist flex" data-shortnr="8" style="line-height: 17px; ">
                                <li class="main">
                                    <div class="ar" style="width: 55px; ">
                                        <b>3x</b>
                                        &nbsp;
                                        <span class="shirt wc w14"></span>
                                    </div>
                                    <div>
                                        <span class="blue"></span>
                                        <a href="https://cx.procyclingstats.com/race/world-championship-cx/2016/result">UCI World Championship</a>
                                        <span style="color: #777; font: 11px tahoma;">&nbsp; ('18, '17, '16)</span>
                                    </div>
                                </li>
                                <li class="main">
                                    <div class="ar" style="width: 55px; ">
                                        <b>4x</b>
                                        &nbsp;
                                    </div>
                                    <div>
                                        <span class="blue"></span>
                                        <a href="https://cx.procyclingstats.com/race/cyclocross-heusden-zolder/2023/result">Telenet Superprestige - Heusden-Zolder</a>
                                        <span style="color: #777; font: 11px tahoma;">&nbsp; ('23, '22, '21, '16)</span>
                                    </div>
                                </li>
                                <li class="main">
                                    <div class="ar" style="width: 55px; ">
                                        <b>2x</b>
                                        &nbsp;
                                    </div>
                                    <div>
                                        <span class="blue"></span>
                                        <a href="https://cx.procyclingstats.com/race/crossvegas/2015/result">Clif Bar Cross Vegas</a>
                                        <span style="color: #777; font: 11px tahoma;">&nbsp; ('16, '15)</span>
                                    </div>
                                </li>
                            </ul>
                        </div>
                    </div>
                    <div class="right w25 mb_w100 mb_mt10">
                        <div class="">
                            <h3>Teams</h3>
                        </div>
                        <div class="">
                            <span class="table-cont">
                                <ul class="list rdr-teams moblist moblist " data-shortnr="4" style=" ">
                                    <li class=" main ">
                                        <div class="season   ">2026</div>
                                        <div class="name   ">
                                            <a href="team/team-visma-lease-a-bike-2026">Team Visma | Lease a Bike</a>
                                        </div>
                                        <div> (WT)</div>
                                        <div class="fs10   "></div>
                                        <div class="clear"></div>
                                    </li>
                                    <li class=" main ">
                                        <div class="season   ">2025</div>
                                        <div class="name   ">
                                            <a href="team/team-visma-lease-a-bike-2025">Team Visma | Lease a Bike</a>
                                        </div>
                                        <div> (WT)</div>
                                        <div class="fs10   "></div>
                                        <div class="clear"></div>
                                    </li>
                                    <li class=" main ">
                                        <div class="season   ">2024</div>
                                        <div class="name   ">
                                            <a href="team/team-visma-lease-a-bike-2024">Team Visma | Lease a Bike</a>
                                        </div>
                                        <div> (WT)</div>
                                        <div class="fs10   "></div>
                                        <div class="clear"></div>
                                    </li>
                                    <li class=" main ">
                                        <div class="season   ">2023</div>
                                        <div class="name   ">
                                            <a href="team/team-jumbo-visma-2023">Jumbo-Visma</a>
                                        </div>
                                        <div> (WT)</div>
                                        <div class="fs10   "></div>
                                        <div class="clear"></div>
                                    </li>
                                    <li class=" main ">
                                        <div class="season   ">2022</div>
                                        <div class="name   ">
                                            <a href="team/team-jumbo-visma-2022">Jumbo-Visma</a>
                                        </div>
                                        <div> (WT)</div>
                                        <div class="fs10   "></div>
                                        <div class="clear"></div>
                                    </li>
                                    <li class=" main ">
                                        <div class="season   ">2021</div>
                                        <div class="name   ">
                                            <a href="team/team-jumbo-visma-2021">Team Jumbo-Visma</a>
                                        </div>
                                        <div> (WT)</div>
                                        <div class="fs10   "></div>
                                        <div class="clear"></div>
                                    </li>
                                    <li class=" main ">
                                        <div class="season   ">2020</div>
                                        <div class="name   ">
                                            <a href="team/team-jumbo-visma-2020">Team Jumbo-Visma</a>
                                        </div>
                                        <div> (WT)</div>
                                        <div class="fs10   "></div>
                                        <div class="clear"></div>
                                    </li>
                                    <li class=" main ">
                                        <div class="season   ">2019</div>
                                        <div class="name   ">
                                            <a href="team/team-jumbo-visma-2019">Team Jumbo-Visma</a>
                                        </div>
                                        <div> (WT)</div>
                                        <div class="fs10   "> as from 01/03</div>
                                        <div class="clear"></div>
                                    </li>
                                    <li class=" main ">
                                        <div class="season   ">2018</div>
                                        <div class="name   ">
                                            <a href="team/verandas-willems-crelan-2018">Vérandas Willems-Crelan </a>
                                        </div>
                                        <div> (PCT)</div>
                                        <div class="fs10   "> until 01/10</div>
                                        <div class="clear"></div>
                                    </li>
                                    <li class=" main ">
                                        <div class="season   ">2017</div>
                                        <div class="name   ">
                                            <a href="team/verandas-willems-crelan-2017">Vérandas Willems-Crelan </a>
                                        </div>
                                        <div> (PCT)</div>
                                        <div class="fs10   "></div>
                                        <div class="clear"></div>
                                    </li>
                                    <li class=" main ">
                                        <div class="season   ">2016</div>
                                        <div class="name   ">
                                            <a href="team/crelan-vastgoedservice-2016">Crelan - Vastgoedservice</a>
                                        </div>
                                        <div> (CT)</div>
                                        <div class="fs10   "></div>
                                        <div class="clear"></div>
                                    </li>
                                    <li class=" main ">
                                        <div class="season   ">2015</div>
                                        <div class="name   ">
                                            <a href="team/pauwels-vastgoedservice-2015">Pauwels - Vastgoedservice</a>
                                        </div>
                                        <div> (CT)</div>
                                        <div class="fs10   "></div>
                                        <div class="clear"></div>
                                    </li>
                                    <li class=" main ">
                                        <div class="season   ">2014</div>
                                        <div class="name   ">
                                            <a href="team/vastgoedservice-golden-palace-2014">Vastgoedservice - Golden Palace</a>
                                        </div>
                                        <div> (CT)</div>
                                        <div class="fs10   "></div>
                                        <div class="clear"></div>
                                    </li>
                                    <li class=" main ">
                                        <div class="season   ">2013</div>
                                        <div class="name   ">
                                            <a href="team/telenet-fidea-2013">Telenet - Fidea</a>
                                        </div>
                                        <div> (CT)</div>
                                        <div class="fs10   "></div>
                                        <div class="clear"></div>
                                    </li>
                                </ul>
                                <a class="showFullList hide showIfMobile" href="">show more</a>
                            </span>
                        </div>
                        <div class="mt30">
                            <h3>Upcoming participations</h3>
                            <span class="table-cont">
                                <ul class="list flex pad2 dashed" style=" ">
                                    <li class="">
                                        <div class="bold   ">24.02</div>
                                        <div class="ellipsis   ">
                                            <span class="flag be"></span>
                                            <a href="race/omloop-het-nieuwsblad/2024/startlist">Omloop Het Nieuwsblad ME</a>
                                        </div>
                                        <div class="fs10 blue   ">
                                            <a href="rider-in-race/wout-van-aert/omloop-het-nieuwsblad">more</a>
                                        </div>
                                        <div class="clear"></div>
                                    </li>
                                    <li class="">
                                        <div class="bold   ">25.02</div>
                                        <div class="ellipsis   ">
                                            <span class="flag be"></span>
                                            <a href="race/kuurne-brussel-kuurne/2024/startlist">Kuurne - Brussel - Kuurne</a>
                                        </div>
                                        <div class="fs10 blue   ">
                                            <a href="rider-in-race/wout-van-aert/kuurne-brussel-kuurne">more</a>
                                        </div>
                                        <div class="clear"></div>
                                    </li>
                                    <li class="">
                                        <div class="bold   ">22.03</div>
                                        <div class="ellipsis   ">
                                            <span class="flag be"></span>
                                            <a href="race/e3-harelbeke/2024/startlist">E3 Saxo Classic</a>
                                        </div>
                                        <div class="fs10 blue   ">
                                            <a href="rider-in-race/wout-van-aert/e3-harelbeke">more</a>
                                        </div>
                                        <div class="clear"></div>
                                    </li>
                                    <li class="">
                                        <div class="bold   ">24.03</div>
                                        <div class="ellipsis   ">
                                            <span class="flag be"></span>
                                            <a href="race/gent-wevelgem/2024/startlist">Gent-Wevelgem in Flanders Fields ME</a>
                                        </div>
                                        <div class="fs10 blue   ">
                                            <a href="rider-in-race/wout-van-aert/gent-wevelgem">more</a>
                                        </div>
                                        <div class="clear"></div>
                                    </li>
                                    <li class="">
                                        <div class="bold   ">27.03</div>
                                        <div class="ellipsis   ">
                                            <span class="flag be"></span>
                                            <a href="race/dwars-door-vlaanderen/2024/startlist">Dwars door Vlaanderen - A travers la Flandre ME</a>
                                        </div>
                                        <div class="fs10 blue   ">
                                            <a href="rider-in-race/wout-van-aert/dwars-door-vlaanderen">more</a>
                                        </div>
                                        <div class="clear"></div>
                                    </li>
                                    <li class="">
                                        <div class="bold   ">31.03</div>
                                        <div class="ellipsis   ">
                                            <span class="flag be"></span>
                                            <a href="race/ronde-van-vlaanderen/2024/startlist">Ronde van Vlaanderen - Tour des Flandres ME</a>
                                        </div>
                                        <div class="fs10 blue   ">
                                            <a href="rider-in-race/wout-van-aert/ronde-van-vlaanderen">more</a>
                                        </div>
                                        <div class="clear"></div>
                                    </li>
                                    <li class="">
                                        <div class="bold   ">07.04</div>
                                        <div class="ellipsis   ">
                                            <span class="flag fr"></span>
                                            <a href="race/paris-roubaix/2024/startlist">Paris-Roubaix</a>
                                        </div>
                                        <div class="fs10 blue   ">
                                            <a href="rider-in-race/wout-van-aert/paris-roubaix">more</a>
                                        </div>
                                        <div class="clear"></div>
                                    </li>
                                    <li class="">
                                        <div class="bold   ">04.05</div>
                                        <div class="ellipsis   ">
                                            <span class="flag it"></span>
                                            <a href="race/giro-d-italia/2024/startlist">Giro d'Italia</a>
                                        </div>
                                        <div class="fs10 blue   ">
                                            <a href="rider-in-race/wout-van-aert/giro-d-italia">more</a>
                                        </div>
                                        <div class="clear"></div>
                                    </li>
                                    <li class="">
                                        <div class="bold   ">03.08</div>
                                        <div class="ellipsis   ">
                                            <span class="flag fr"></span>
                                            <a href="race/olympic-games/2024/startlist">Olympic Games ME - Road Race</a>
                                        </div>
                                        <div class="fs10 blue   ">
                                            <a href="rider-in-race/wout-van-aert/olympic-games">more</a>
                                        </div>
                                        <div class="clear"></div>
                                    </li>
                                    <li class="">
                                        <div class="bold   ">17.08</div>
                                        <div class="ellipsis   ">
                                            <span class="flag es"></span>
                                            <a href="race/vuelta-a-espana/2024/startlist">La Vuelta Ciclista a España</a>
                                        </div>
                                        <div class="fs10 blue   ">
                                            <a href="rider-in-race/wout-van-aert/vuelta-a-espana">more</a>
                                        </div>
                                        <div class="clear"></div>
                                    </li>
                                </ul>
                            </span>
                        </div>
                        <div class="">
                            <div class="">
                                <div class="mob-ad rider-after-program"></div>
                            </div>
                        </div>
                    </div>
                    <div class="left w73 mb_w100 mt20" style="min-height: 400px; ">
                        <ul class="tabsV3" style=" ">
                            <li class="cur">
                                <a class="toggleDiscipline" data-discipline="road" data-rid="168961" data-season="2024" href="">road</a>
                                <span></span>
                            </li>
                            <li class="">
                                <a class="toggleDiscipline" data-discipline="cx" data-rid="1000025" data-season="2024" href="">cross</a>
                                <span></span>
                            </li>
                        </ul>
                        <ul class="rdrSeasonNav " data-discipline="road">
                            <li data-season="2024" class="cur">
                                <a class="seasonResults" data-season="2024" href="">2024</a>
                            </li>
                            <li data-season="2023" class="">
                                <a class="seasonResults" data-season="2023" href="">2023</a>
                            </li>
                            <li data-season="2022" class="">
                                <a class="seasonResults" data-season="2022" href="">2022</a>
                            </li>
                            <li data-season="2021" class="">
                                <a class="seasonResults" data-season="2021" href="">2021</a>
                            </li>
                            <li data-season="2020" class="">
                                <a class="seasonResults" data-season="2020" href="">2020</a>
                            </li>
                            <li data-season="2019" class="">
                                <a class="seasonResults" data-season="2019" href="">2019</a>
                            </li>
                            <li data-season="2018" class="">
                                <a class="seasonResults" data-season="2018" href="">2018</a>
                            </li>
                            <li data-season="2017" class="">
                                <a class="seasonResults" data-season="2017" href="">2017</a>
                            </li>
                            <li data-season="2016" class="">
                                <a class="seasonResults" data-season="2016" href="">2016</a>
                            </li>
                            <li data-season="2015" class="">
                                <a class="seasonResults" data-season="2015" href="">2015</a>
                            </li>
                            <li class="more  show_if">
                                <span>more</span>
                                <ul>
                                    <li data-season="2024" class=" ">
                                        <a class="seasonResults" data-season="2024" href="">2024</a>
                                    </li>
                                    <li data-season="2023" class=" ">
                                        <a class="seasonResults" data-season="2023" href="">2023</a>
                                    </li>
                                    <li data-season="2022" class=" ">
                                        <a class="seasonResults" data-season="2022" href="">2022</a>
                                    </li>
                                    <li data-season="2021" class=" ">
                                        <a class="seasonResults" data-season="2021" href="">2021</a>
                                    </li>
                                    <li data-season="2020" class=" ">
                                        <a class="seasonResults" data-season="2020" href="">2020</a>
                                    </li>
                                    <li data-season="2019" class=" ">
                                        <a class="seasonResults" data-season="2019" href="">2019</a>
                                    </li>
                                    <li data-season="2018" class=" ">
                                        <a class="seasonResults" data-season="2018" href="">2018</a>
                                    </li>
                                    <li data-season="2017" class=" ">
                                        <a class="seasonResults" data-season="2017" href="">2017</a>
                                    </li>
                                    <li data-season="2016" class=" ">
                                        <a class="seasonResults" data-season="2016" href="">2016</a>
                                    </li>
                                    <li data-season="2015" class=" ">
                                        <a class="seasonResults" data-season="2015" href="">2015</a>
                                    </li>
                                    <li data-season="2014" class=" ">
                                        <a class="seasonResults" data-season="2014" href="">2014</a>
                                    </li>
                                    <li data-season="2013" class=" ">
                                        <a class="seasonResults" data-season="2013" href="">2013</a>
                                    </li>
                                    <li data-season="2012" class=" ">
                                        <a class="seasonResults" data-season="2012" href="">2012</a>
                                    </li>
                                </ul>
                            </li>
                        </ul>
                        <ul class="rdrSeasonNav hide" data-discipline="cx">
                            <li data-season="2024" class="cur">
                                <a class="seasonResults" data-season="2024" href="">23/24</a>
                            </li>
                            <li data-season="2023" class="">
                                <a class="seasonResults" data-season="2023" href="">22/23</a>
                            </li>
                            <li data-season="2022" class="">
                                <a class="seasonResults" data-season="2022" href="">21/22</a>
                            </li>
                            <li data-season="2021" class="">
                                <a class="seasonResults" data-season="2021" href="">20/21</a>
                            </li>
                            <li data-season="2020" class="">
                                <a class="seasonResults" data-season="2020" href="">19/20</a>
                            </li>
                            <li data-season="2019" class="">
                                <a class="seasonResults" data-season="2019" href="">18/19</a>
                            </li>
                            <li data-season="2018" class="">
                                <a class="seasonResults" data-season="2018" href="">17/18</a>
                            </li>
                            <li data-season="2017" class="">
                                <a class="seasonResults" data-season="2017" href="">16/17</a>
                            </li>
                            <li data-season="2016" class="">
                                <a class="seasonResults" data-season="2016" href="">15/16</a>
                            </li>
                            <li data-season="2015" class="">
                                <a class="seasonResults" data-season="2015" href="">14/15</a>
                            </li>
                            <li class="more  show_if">
                                <span>more</span>
                                <ul>
                                    <li data-season="2024" class=" ">
                                        <a class="seasonResults" data-season="2024" href="">23/24</a>
                                    </li>
                                    <li data-season="2023" class=" ">
                                        <a class="seasonResults" data-season="2023" href="">22/23</a>
                                    </li>
                                    <li data-season="2022" class=" ">
                                        <a class="seasonResults" data-season="2022" href="">21/22</a>
                                    </li>
                                    <li data-season="2021" class=" ">
                                        <a class="seasonResults" data-season="2021" href="">20/21</a>
                                    </li>
                                    <li data-season="2020" class=" ">
                                        <a class="seasonResults" data-season="2020" href="">19/20</a>
                                    </li>
                                    <li data-season="2019" class=" ">
                                        <a class="seasonResults" data-season="2019" href="">18/19</a>
                                    </li>
                                    <li data-season="2018" class=" ">
                                        <a class="seasonResults" data-season="2018" href="">17/18</a>
                                    </li>
                                    <li data-season="2017" class=" ">
                                        <a class="seasonResults" data-season="2017" href="">16/17</a>
                                    </li>
                                    <li data-season="2016" class=" ">
                                        <a class="seasonResults" data-season="2016" href="">15/16</a>
                                    </li>
                                    <li data-season="2015" class=" ">
                                        <a class="seasonResults" data-season="2015" href="">14/15</a>
                                    </li>
                                    <li data-season="2014" class=" ">
                                        <a class="seasonResults" data-season="2014" href="">13/14</a>
                                    </li>
                                    <li data-season="2013" class=" ">
                                        <a class="seasonResults" data-season="2013" href="">12/13</a>
                                    </li>
                                    <li data-season="2012" class=" ">
                                        <a class="seasonResults" data-season="2012" href="">11/12</a>
                                    </li>
                                    <li data-season="2011" class=" ">
                                        <a class="seasonResults" data-season="2011" href="">10/11</a>
                                    </li>
                                </ul>
                            </li>
                        </ul>
                        <div class="sortByResults">
                            Sort by: 
                            <a class=" " data-case="date" href="">Date</a>
                            <a class=" " data-case="rnk" href="">Result</a>
                            <a class=" " data-case="pnt" href="">Points</a>
                        </div>
                        <div class="clear"></div>
                        <div id="resultsCont">
                            <table class="rdrResults">
                                <thead>
                                    <tr>
                                        <th style="width: 11%;">
                                            <br/>
                                            Date
                                        </th>
                                        <th style="width: 4%;">
                                            <br/>
                                            Result
                                        </th>
                                        <th class="cu600" style="width: 4%;"></th>
                                        <th style="width: 2%;"></th>
                                        <th class="name" style="">
                                            <br/>
                                            Race
                                        </th>
                                        <th class="cu600" style="width: 8%;">
                                            <br/>
                                            Distance
                                        </th>
                                        <th class="cu600" style="width: 8%;">
                                            Points
                                            <br/>
                                            PCS
                                        </th>
                                        <th class="cu600 " style="width: 8%;">
                                            Points
                                            <br/>
                                            UCI
                                        </th>
                                        <th class="cu600" style="width: 3%;"></th>
                                    </tr>
                                </thead>
                                <tbody>
                                    <tr data-main="1">
                                        <td>14.02 &raquo; 18.02</td>
                                        <td></td>
                                        <td class="gc cu600"></td>
                                        <td style="padding: 0; ">
                                            <span class=""></span>
                                        </td>
                                        <td class="name">
                                            <span class="flag pt"></span>
                                            <a href="race/volta-ao-algarve/2024/stage-5-kom">
                                                Volta ao Algarve em Bicicleta 
                                                <span>(2.Pro)</span>
                                            </a>
                                        </td>
                                        <td class="cu600"></td>
                                        <td class="cu600"></td>
                                        <td class="cu600 "></td>
                                        <td class="cu600">
                                            <a class="more" href="rider-in-race/wout-van-aert/volta-ao-algarve">more</a>
                                        </td>
                                    </tr>
                                    <tr data-main="0">
                                        <td></td>
                                        <td>6</td>
                                        <td class="gc cu600"></td>
                                        <td style="padding: 0; ">
                                            <span class="icon shirt w14 st7"></span>
                                        </td>
                                        <td class="name">
                                            <span class="flag pt"></span>
                                            <a href="race/volta-ao-algarve/2024/stage-5-kom">Mountains classification</a>
                                        </td>
                                        <td class="cu600"></td>
                                        <td class="cu600"></td>
                                        <td class="cu600 "></td>
                                        <td class="cu600">
                                            <a class="more" href="rider-in-race/wout-van-aert/volta-ao-algarve">more</a>
                                        </td>
                                    </tr>
                                    <tr data-main="0">
                                        <td></td>
                                        <td>3</td>
                                        <td class="gc cu600"></td>
                                        <td style="padding: 0; ">
                                            <span class="icon shirt w14 st5"></span>
                                        </td>
                                        <td class="name">
                                            <span class="flag pt"></span>
                                            <a href="race/volta-ao-algarve/2024/stage-5-points">Points classification</a>
                                        </td>
                                        <td class="cu600"></td>
                                        <td class="cu600"></td>
                                        <td class="cu600 "></td>
                                        <td class="cu600">
                                            <a class="more" href="rider-in-race/wout-van-aert/volta-ao-algarve">more</a>
                                        </td>
                                    </tr>
                                    <tr data-main="0">
                                        <td></td>
                                        <td>7</td>
                                        <td class="gc cu600"></td>
                                        <td style="padding: 0; ">
                                            <span class="icon shirt w14 st4"></span>
                                        </td>
                                        <td class="name">
                                            <span class="flag pt"></span>
                                            <a href="race/volta-ao-algarve/2024/gc">General classification</a>
                                        </td>
                                        <td class="cu600"></td>
                                        <td class="cu600">55</td>
                                        <td class="cu600 ">60</td>
                                        <td class="cu600">
                                            <a class="more" href="rider-in-race/wout-van-aert/volta-ao-algarve">more</a>
                                        </td>
                                    </tr>
                                    <tr data-main="0">
                                        <td>18.02</td>
                                        <td>13</td>
                                        <td class="gc cu600"></td>
                                        <td style="padding: 0; ">
                                            <span class=""></span>
                                        </td>
                                        <td class="name">
                                            <span class="flag pt"></span>
                                            <a href="race/volta-ao-algarve/2024/stage-5">Stage 5 - Faro &rsaquo; Alto do Malhāo</a>
                                        </td>
                                        <td class="cu600">165.8</td>
                                        <td class="cu600"></td>
                                        <td class="cu600 "></td>
                                        <td class="cu600">
                                            <a class="more" href="rider-in-race/wout-van-aert/volta-ao-algarve">more</a>
                                        </td>
                                    </tr>
                                    <tr data-main="0">
                                        <td>17.02</td>
                                        <td>11</td>
                                        <td class="gc cu600">4</td>
                                        <td style="padding: 0; ">
                                            <span class="icon chrono"></span>
                                        </td>
                                        <td class="name">
                                            <span class="flag pt"></span>
                                            <a href="race/volta-ao-algarve/2024/stage-4">Stage 4 (ITT) - Albufeira &rsaquo; Albufeira</a>
                                        </td>
                                        <td class="cu600">22</td>
                                        <td class="cu600"></td>
                                        <td class="cu600 "></td>
                                        <td class="cu600">
                                            <a class="more" href="rider-in-race/wout-van-aert/volta-ao-algarve">more</a>
                                        </td>
                                    </tr>
                                    <tr data-main="0">
                                        <td>16.02</td>
                                        <td>1</td>
                                        <td class="gc cu600">8</td>
                                        <td style="padding: 0; ">
                                            <span class=""></span>
                                        </td>
                                        <td class="name">
                                            <span class="flag pt"></span>
                                            <a href="race/volta-ao-algarve/2024/stage-3">Stage 3 - Vila Real de Santo António &rsaquo; Tavira</a>
                                        </td>
                                        <td class="cu600">192.2</td>
                                        <td class="cu600">30</td>
                                        <td class="cu600 ">20</td>
                                        <td class="cu600">
                                            <a class="more" href="rider-in-race/wout-van-aert/volta-ao-algarve">more</a>
                                        </td>
                                    </tr>
                                    <tr data-main="0">
                                        <td>15.02</td>
                                        <td>16</td>
                                        <td class="gc cu600">14</td>
                                        <td style="padding: 0; ">
                                            <span class=""></span>
                                        </td>
                                        <td class="name">
                                            <span class="flag pt"></span>
                                            <a href="race/volta-ao-algarve/2024/stage-2">Stage 2 - Lagoa &rsaquo; Alto da Fóia</a>
                                        </td>
                                        <td class="cu600">171.9</td>
                                        <td class="cu600"></td>
                                        <td class="cu600 "></td>
                                        <td class="cu600">
                                            <a class="more" href="rider-in-race/wout-van-aert/volta-ao-algarve">more</a>
                                        </td>
                                    </tr>
                                    <tr data-main="0">
                                        <td>14.02</td>
                                        <td>135</td>
                                        <td class="gc cu600">136</td>
                                        <td style="padding: 0; ">
                                            <span class=""></span>
                                        </td>
                                        <td class="name">
                                            <span class="flag pt"></span>
                                            <a href="race/volta-ao-algarve/2024/stage-1">Stage 1 - Portimāo &rsaquo; Lagos</a>
                                        </td>
                                        <td class="cu600">200.8</td>
                                        <td class="cu600"></td>
                                        <td class="cu600 "></td>
                                        <td class="cu600">
                                            <a class="more" href="rider-in-race/wout-van-aert/volta-ao-algarve">more</a>
                                        </td>
                                    </tr>
                                    <tr data-main="1">
                                        <td>12.02</td>
                                        <td>45</td>
                                        <td class="gc cu600"></td>
                                        <td style="padding: 0; ">
                                            <span class=""></span>
                                        </td>
                                        <td class="name">
                                            <span class="flag es"></span>
                                            <a href="race/clasica-jaen-paraiso-interior/2024/result">
                                                Clásica Jaén 
                                                <span>(1.1)</span>
                                            </a>
                                        </td>
                                        <td class="cu600">158.3</td>
                                        <td class="cu600"></td>
                                        <td class="cu600 "></td>
                                        <td class="cu600">
                                            <a class="more" href="rider-in-race/wout-van-aert/clasica-jaen-paraiso-interior">more</a>
                                        </td>
                                    </tr>
                                    <tr data-main="1">
                                        <td>11.02</td>
                                        <td>10</td>
                                        <td class="gc cu600"></td>
                                        <td style="padding: 0; ">
                                            <span class=""></span>
                                        </td>
                                        <td class="name">
                                            <span class="flag es"></span>
                                            <a href="race/clasica-de-almeria/2024/result">
                                                Clasica de Almeria 
                                                <span>(1.Pro)</span>
                                            </a>
                                        </td>
                                        <td class="cu600">192.3</td>
                                        <td class="cu600">22</td>
                                        <td class="cu600 ">35</td>
                                        <td class="cu600">
                                            <a class="more" href="rider-in-race/wout-van-aert/clasica-de-almeria">more</a>
                                        </td>
                                    </tr>
                                </tbody>
                            </table>
                            <div class="rdrResultsSum">
                                <div>
                                    1103</b>
                                     km in 
                                    <b>7</b>
                                     days | PCS points: 
                                    <b>107</b>
                                     |  UCI points: 
                                    <b>115</b>
                                </div>
                            </div>
                        </div>
                        <div style="float: right; margin-top: 8px; ">
                            <span style=" font-size: 10px; color: #aaa; letter-spacing: -0.5px; margin-left: 10px;">
                                <svg viewBox="0 0 24 24" style="width: 10px; ">
                                    <g>
                                        <path d="M0 0h24v24H0z" fill="none"/>
                                        <path fill="#f55733" d="M3.783 2.826L12 1l8.217 1.826a1 1 0 0 1 .783.976v9.987a6 6 0 0 1-2.672 4.992L12 23l-6.328-4.219A6 6 0 0 1 3 13.79V3.802a1 1 0 0 1 .783-.976z"/>
                                    </g>
                                </svg>
                                 = number of kilometres in a group before the peloton
                            </span>
                        </div>
                        <div class="clear"></div>
                    </div>
                    <div class="right w25 mb_w100">
                        <div class="mt30">
                            <h3>Key statistics</h3>
                            <span class="table-cont">
                                <ul class="rider-kpi" style=" ">
                                    <li class="">
                                        <div class="nr   ">45</div>
                                        <div class="title   ">
                                            <a href="rider/wout-van-aert/statistics/wins">Wins</a>
                                        </div>
                                        <div class="info fs11   ">GC (3)</div>
                                        <div class="info fs11   ">Oneday races (17)</div>
                                        <div class="info fs11   ">ITT (9)</div>
                                        <div class="info fs11   "></div>
                                        <div class="info fs11   "></div>
                                        <div class="clear"></div>
                                    </li>
                                    <li class="">
                                        <div class="nr   ">5</div>
                                        <div class="title   ">
                                            <a href="rider/wout-van-aert/statistics/grand-tour-starts">Grand tours</a>
                                        </div>
                                        <div class="info fs11   ">tour (5) giro (0) vuelta(0)</div>
                                        <div class="info fs11   "></div>
                                        <div class="info fs11   "></div>
                                        <div class="info fs11   "></div>
                                        <div class="info fs11   "></div>
                                        <div class="clear"></div>
                                    </li>
                                    <li class="">
                                        <div class="nr   ">16</div>
                                        <div class="title   ">
                                            <a href="rider/wout-van-aert/statistics/top-classic-results">Classics</a>
                                        </div>
                                        <div class="info fs11   ">RBX(5)</div>
                                        <div class="info fs11   ">MSR(5)</div>
                                        <div class="info fs11   ">RVV(5)</div>
                                        <div class="info fs11   ">LBL(1)</div>
                                        <div class="info fs11   ">LOM(0)</div>
                                        <div class="clear"></div>
                                    </li>
                                </ul>
                            </span>
                        </div>
                        <div class="mt30">
                            <h3>PCS Ranking position per season</h3>
                            <span class="table-cont">
                                <table class="basic rdr-season-stats" style="">
                                    <thead>
                                        <tr>
                                            <th class="season  " style="width: 16.7%; "></th>
                                            <th class="bar  " style="width: 66.7%; ">Points</th>
                                            <th class="ac  " style="width: 16.7%; ">position</th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                        <tr>
                                            <td class="season  ">2024</td>
                                            <td class="bar  ">
                                                <div class="barCont">
                                                    <div class="bg left green2 h12" style="width: 3%; position: relative;   ">
                                                        <div class="vTitle" style="line-height: 12px;">107</div>
                                                    </div>
                                                </div>
                                            </td>
                                            <td class="ac  ">49</td>
                                        </tr>
                                        <tr>
                                            <td class="season  ">2023</td>
                                            <td class="bar  ">
                                                <div class="barCont">
                                                    <div class="bg left green2 h12" style="width: 57.6%; position: relative;   ">
                                                        <div class="vTitle" style="line-height: 12px;">2060</div>
                                                    </div>
                                                </div>
                                            </td>
                                            <td class="ac  ">6</td>
                                        </tr>
                                        <tr>
                                            <td class="season  ">2022</td>
                                            <td class="bar  ">
                                                <div class="barCont">
                                                    <div class="bg left green2 h12" style="width: 80%; position: relative;   ">
                                                        <div class="vTitle" style="line-height: 12px;">2862</div>
                                                    </div>
                                                </div>
                                            </td>
                                            <td class="ac  ">2</td>
                                        </tr>
                                        <tr>
                                            <td class="season  ">2021</td>
                                            <td class="bar  ">
                                                <div class="barCont">
                                                    <div class="bg left green2 h12" style="width: 78.6%; position: relative;   ">
                                                        <div class="vTitle" style="line-height: 12px;">2813</div>
                                                    </div>
                                                </div>
                                            </td>
                                            <td class="ac  ">2</td>
                                        </tr>
                                        <tr>
                                            <td class="season  ">2020</td>
                                            <td class="bar  ">
                                                <div class="barCont">
                                                    <div class="bg left green2 h12" style="width: 51.3%; position: relative;   ">
                                                        <div class="vTitle" style="line-height: 12px;">1834</div>
                                                    </div>
                                                </div>
                                            </td>
                                            <td class="ac  ">3</td>
                                        </tr>
                                        <tr>
                                            <td class="season  ">2019</td>
                                            <td class="bar  ">
                                                <div class="barCont">
                                                    <div class="bg left green2 h12" style="width: 23.8%; position: relative;   ">
                                                        <div class="vTitle" style="line-height: 12px;">852</div>
                                                    </div>
                                                </div>
                                            </td>
                                            <td class="ac  ">53</td>
                                        </tr>
                                        <tr>
                                            <td class="season  ">2018</td>
                                            <td class="bar  ">
                                                <div class="barCont">
                                                    <div class="bg left green2 h12" style="width: 15.4%; position: relative;   ">
                                                        <div class="vTitle" style="line-height: 12px;">551</div>
                                                    </div>
                                                </div>
                                            </td>
                                            <td class="ac  ">95</td>
                                        </tr>
                                        <tr>
                                            <td class="season  ">2017</td>
                                            <td class="bar  ">
                                                <div class="barCont">
                                                    <div class="bg left green2 h12" style="width: 12.7%; position: relative;   ">
                                                        <div class="vTitle" style="line-height: 12px;">456</div>
                                                    </div>
                                                </div>
                                            </td>
                                            <td class="ac  ">116</td>
                                        </tr>
                                        <tr>
                                            <td class="season  ">2016</td>
                                            <td class="bar  ">
                                                <div class="barCont">
                                                    <div class="bg left green2 h12" style="width: 7.1%; position: relative;   ">
                                                        <div class="vTitle" style="line-height: 12px;">254</div>
                                                    </div>
                                                </div>
                                            </td>
                                            <td class="ac  ">247</td>
                                        </tr>
                                        <tr>
                                            <td class="season  ">2015</td>
                                            <td class="bar  ">
                                                <div class="barCont">
                                                    <div class="bg left green2 h12" style="width: 1.2%; position: relative;   ">
                                                        <div class="vTitle" style="line-height: 12px;">43</div>
                                                    </div>
                                                </div>
                                            </td>
                                            <td class="ac  ">936</td>
                                        </tr>
                                        <tr>
                                            <td class="season  ">2014</td>
                                            <td class="bar  ">
                                                <div class="barCont">
                                                    <div class="bg left green2 h12" style="width: 0.5%; position: relative;   ">
                                                        <div class="vTitle" style="line-height: 12px;">18</div>
                                                    </div>
                                                </div>
                                            </td>
                                            <td class="ac  ">1312</td>
                                        </tr>
                                    </tbody>
                                </table>
                            </span>
                        </div>
                        <div class="mt30">
                            <h3>Family</h3>
                            <span class="table-cont">
                                <ul class="list" style=" ">
                                    <li class="">
                                        <div>
                                            <a href="rider/jos-van-aert">Jos van Aert</a>
                                        </div>
                                        <div> (Uncle)</div>
                                        <div class="clear"></div>
                                    </li>
                                </ul>
                            </span>
                        </div>
                        <div class="mt30">
                            <h3>H2H Suggestions</h3>
                            <span class="table-cont">
                                <ul class="list" style=" ">
                                    <li class="">
                                        <div>
                                            <a href="rider-vs-rider/wout-van-aert/mathieu-van-der-poel">VAN DER POEL Mathieu</a>
                                        </div>
                                        <div class="clear"></div>
                                    </li>
                                    <li class="">
                                        <div>
                                            <a href="rider-vs-rider/wout-van-aert/peter-sagan">SAGAN Peter</a>
                                        </div>
                                        <div class="clear"></div>
                                    </li>
                                    <li class="">
                                        <div>
                                            <a href="rider-vs-rider/wout-van-aert/julian-alaphilippe">ALAPHILIPPE Julian</a>
                                        </div>
                                        <div class="clear"></div>
                                    </li>
                                    <li class="">
                                        <div>
                                            <a href="rider-vs-rider/wout-van-aert/eddy-merckx">MERCKX Eddy</a>
                                        </div>
                                        <div class="clear"></div>
                                    </li>
                                    <li class="">
                                        <div>
                                            <a href="rider-vs-rider/wout-van-aert/roger-de-vlaeminck">DE VLAEMINCK Roger</a>
                                        </div>
                                        <div class="clear"></div>
                                    </li>
                                </ul>
                            </span>
                        </div>
                    </div>
                    <div class="clear"></div>
                    <script>
                    $(document).ready(function() {
                        var id = 168961;
                        var natid = 3012724;
                        var cxid = 1000025;
                        var seo = 'wout-van-aert';
                        var sort = 'date';
                        var pid = '1238';
                        var discipline = 'road';
                        var subpath = '';

                        $(document).on('click', '.toggleDiscipline', function(e) {
                            e.preventDefault();
                            discipline = $(this).attr('data-discipline');
                            id = $(this).attr('data-rid');
                            console.log('dis: ' + discipline + ', id: ' + $(this).attr('data-rid') + ', season: ' + $(this).attr('data-season'));
                            $('.rdrSeasonNav').hide();
                            $('.rdrSeasonNav[data-discipline="' + discipline + '"]').show();
                            $(".rdrSeasonNav li").removeClass('cur');
                            $(".rdrSeasonNav li[data-season=" + $(this).attr('data-season') + "]").addClass('cur');
                            var url = 'rider/' + seo + '/' + discipline;
                            //history.pushState("string", "page", url);  
                            $(this).parents('ul').find('li').removeClass('cur');
                            $(this).parent('li').addClass('cur');
                            //console.log(subpath);
                            $.ajax({
                                method: "POST",
                                url: subpath + "rdr/start3.php",
                                data: {
                                    getresults: 1,
                                    id: id,
                                    natid: natid,
                                    cxid: cxid,
                                    sort: sort,
                                    pid: pid,
                                    discipline: discipline,
                                    season: $(this).attr('data-season')
                                }
                            }).done(function(response) {
                                $('#resultsCont').html(response);
                            })
                        });

                        $(document).on('click', '.seasonResults', function(e) {
                            e.preventDefault();
                            var season = $(this).attr('data-season');


                            $("#resultsCont").html('<br /><div class=\"loading\"> loading results ..</div>');
                            var url = 'rider/' + seo + '/' + season;
                            history.pushState("string", "page", url);
                            $(".rdrSeasonNav li").removeClass('cur');
                            $(".rdrSeasonNav li[data-season=" + season + "]").addClass('cur');

                            $.ajax({
                                method: "POST",
                                url: subpath + "rdr/start3.php",
                                data: {
                                    getresults: 1,
                                    id: id,
                                    natid: natid,
                                    cxid: cxid,
                                    season: season,
                                    sort: sort,
                                    pid: pid,
                                    discipline: discipline
                                }
                            }).done(function(response) {
                                //console.log(response);
                                $('#resultsCont').html(response);
                            })

                        });

                        $(".sortByResults a").click(function(e) {
                            e.preventDefault();
                            var season = $(".rdrSeasonNav li.cur").attr('data-season');
                            $("#resultsCont").html('<br /><div class=\"loading\"> loading results ..</div>');
                            $(".sortByResults a").removeClass('cur');
                            $(this).addClass('cur');
                            sort = $(this).attr('data-case');
                            // console.log('dis: '+discipline+', season: '+season+', sort: '+sort+', cxid: '+cxid+', natid: '+natid+', id: '+id);

                            $.ajax({
                                method: "POST",
                                url: subpath + "rdr/start3.php",
                                data: {
                                    getresults: 1,
                                    id: id,
                                    natid: natid,
                                    cxid: cxid,
                                    season: season,
                                    sort: $(this).attr('data-case'),
                                    pid: pid,
                                    discipline: discipline
                                }
                            }).done(function(response) {
                                $('#resultsCont').html(response);
                            })
                        });

                    });
                    </script>

                </div>
                <div class="desktop-ad below-content rider"></div>
                <div class="mob-ad below-content rider"></div>
            </div>
            <div style="clear: both; "></div>
        </div>
        </center>
        <div class="push"></div>
    </div>
    <div class="footer">
        <div class="sfa">
            <ul class="ftr-list">
                <li class="group">
                    <h3>Grand Tours</h3>
                    <ul>
                        <li>
                            <a href="race/tour-de-france">Tour de France</a>
                        </li>
                        <li>
                            <a href="race/giro-d-italia">Giro d'Italia</a>
                        </li>
                        <li>
                            <a href="race/vuelta-a-espana">Vuelta a España</a>
                        </li>
                    </ul>
                    <h3>Major Tours</h3>
                    <ul>
                        <li>
                            <a href="race/paris-nice">Paris-Nice</a>
                        </li>
                        <li>
                            <a href="race/tirreno-adriatico">Tirreno-Adriatico</a>
                        </li>
                        <li>
                            <a href="race/volta-a-catalunya">Volta a Catalunya</a>
                        </li>
                        <li>
                            <a href="race/tour-de-romandie">Tour de Romandie</a>
                        </li>
                        <li>
                            <a href="race/tour-de-suisse">Tour de Suisse</a>
                        </li>
                        <li>
                            <a href="race/dauphine">Critérium du Dauphiné</a>
                        </li>
                        <li>
                            <a href="race/itzulia-basque-country">Itzulia Basque Country</a>
                        </li>
                    </ul>
                </li>
                <li class="group">
                    <h3>Monuments</h3>
                    <ul>
                        <li>
                            <a href="race/milano-sanremo">Milano-SanRemo</a>
                        </li>
                        <li>
                            <a href="race/ronde-van-vlaanderen">Ronde van Vlaanderen</a>
                        </li>
                        <li>
                            <a href="race/paris-roubaix">Paris-Roubaix</a>
                        </li>
                        <li>
                            <a href="race/liege-bastogne-liege">Liège-Bastogne-Liège</a>
                        </li>
                        <li>
                            <a href="race/il-lombardia">Il Lombardia</a>
                        </li>
                    </ul>
                    <h3>Championships</h3>
                    <ul>
                        <li>
                            <a href="race/world-championship">World Championships</a>
                        </li>
                        <li>
                            <a href="race/uec-road-european-championships">European championships</a>
                        </li>
                    </ul>
                </li>
                <li class="group">
                    <h3>Top classics</h3>
                    <ul>
                        <li>
                            <a href="race/omloop-het-nieuwsblad">Omloop Het Nieuwsblad</a>
                        </li>
                        <li>
                            <a href="race/strade-bianche">Strade Bianche</a>
                        </li>
                        <li>
                            <a href="race/e3-harelbeke">E3 Classic</a>
                        </li>
                        <li>
                            <a href="race/gent-wevelgem">Gent-Wevelgem</a>
                        </li>
                        <li>
                            <a href="race/dwars-door-vlaanderen">Dwars door vlaanderen</a>
                        </li>
                        <li>
                            <a href="race/eschborn-frankfurt">Eschborn-Frankfurt</a>
                        </li>
                        <li>
                            <a href="race/amstel-gold-race">Amstel Gold Race</a>
                        </li>
                        <li>
                            <a href="race/la-fleche-wallone">La Flèche Wallonne</a>
                        </li>
                        <li>
                            <a href="race/san-sebastian">San Sebastian</a>
                        </li>
                        <li>
                            <a href="race/bretagne-classic">Bretagne Classic</a>
                        </li>
                        <li>
                            <a href="race/gp-quebec">GP Québec</a>
                        </li>
                        <li>
                            <a href="race/gp-montreal">GP Montréal</a>
                        </li>
                    </ul>
                </li>
                <li class="group">
                    <h3>Popular riders</h3>
                    <ul>
                        <li>
                            <a href="rider/tadej-pogacar">Tadej Pogačar</a>
                        </li>
                        <li>
                            <a href="rider/wout-van-aert">Wout van Aert</a>
                        </li>
                        <li>
                            <a href="rider/remco-evenepoel">Remco Evenepoel</a>
                        </li>
                        <li>
                            <a href="rider/jonas-vingegaard-rasmussen">Jonas Vingegaard</a>
                        </li>
                        <li>
                            <a href="rider/mathieu-van-der-poel">Mathieu van der Poel</a>
                        </li>
                        <li>
                            <a href="rider/mads-pedersen">Mads Pedersen</a>
                        </li>
                        <li>
                            <a href="rider/primoz-roglic">Primoz Roglic</a>
                        </li>
                        <li>
                            <a href="rider/demi-vollering">Demi Vollering</a>
                        </li>
                        <li>
                            <a href="rider/lotte-kopecky">Lotte Kopecky</a>
                        </li>
                        <li>
                            <a href="rider/katarzyna-niewiadoma">Katarzyna Niewiadoma</a>
                        </li>
                    </ul>
                </li>
                <li class="group">
                    <h3>Rankings</h3>
                    <ul>
                        <li>
                            <a href="rankings.php">PCS ranking</a>
                        </li>
                        <li>
                            <a href="rankings/me/uci-individual">UCI World Ranking</a>
                        </li>
                        <li>
                            <a href="rankings/me/all-time">Alltime</a>
                        </li>
                    </ul>
                    <h3>Statistics</h3>
                    <ul>
                        <li>
                            <a href="statistics/start/points-per-age">Points per age</a>
                        </li>
                        <li>
                            <a href="statistics/start/latest-injuries">Latest injuries</a>
                        </li>
                        <li>
                            <a href="statistics/riders/youngest-riders">Youngest riders</a>
                        </li>
                        <li>
                            <a href="statistics/grandtours">Grand tour statistics</a>
                        </li>
                        <li>
                            <a href="statistics/monuments">Monument classics</a>
                        </li>
                        <li>
                            <a href="teams/transfers">Latest transfers</a>
                        </li>
                        <li>
                            <a href="favorite500">Favorite 500</a>
                        </li>
                    </ul>
                </li>
                <li class="group">
                    <h3>Info</h3>
                    <ul>
                        <li>
                            <a href="info/point-scales">Points scales</a>
                        </li>
                        <li>
                            <a href="info/profile-score-explained">Profile scores</a>
                        </li>
                        <li>
                            <a href="reset_password.php">Reset password</a>
                        </li>
                        <li>
                            <a href="info/pcs-game">PCS game</a>
                        </li>
                        <li>
                            <a href="info/procyclinggame">ProCyclingGame</a>
                        </li>
                    </ul>
                    <h3>About ProCyclingStats</h3>
                    <ul>
                        <li>
                            <a href="info/about-us">About us</a>
                        </li>
                        <li>
                            <a href="info/contact">Contact us</a>
                        </li>
                        <li>
                            <a href="info/cookie-policy">Cookie policy</a>
                        </li>
                        <li>
                            <a href="info/sitemap">Sitemap</a>
                        </li>
                        <li>
                            <a href="info/contributions">Contributions</a>
                        </li>
                        <li>
                            <span>Pageload 0.0963s</span>
                        </li>
                    </ul>
                </li>
            </ul>
            <input type="hidden" id="pageid" value="1238"/>
            <input type="hidden" id="userid" value="0"/>
        </div>
    </div>

</body>
</html>
<script src="search_list9.js"></script>
`

func TestUpcomingRaces(t *testing.T) {
	stats, err := Extract(content)
	if err != nil {
		t.Fatal(err)
	}
	if len(stats.UpcomingRaces) != 10 {
		t.Fatalf("expected 10 upcoming races, got %v", stats.UpcomingRaces)
	}
}

func TestResults(t *testing.T) {
	stats, err := Extract(content)
	if err != nil {
		t.Fatal(err)
	}
	if len(stats.Results) != 12 {
		t.Fatalf("expected 12 results, got %v", stats.Results)
	}
}
