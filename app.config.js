'use strict';



var appConfig = window.appConfig || {};

appConfig.menu_speed = 200;

appConfig.smartSkin = "smart-style-0";

appConfig.skins = [{
        name: "smart-style-0",
        logo: "styles/img/logo.png",
        class: "btn btn-block btn-xs txt-color-white margin-right-5",
        style: "background-color:#4E463F;",
        label: "Smart Default"
    },

    {
        name: "smart-style-1",
        logo: "styles/img/logo-white.png",
        class: "btn btn-block btn-xs txt-color-white",
        style: "background:#3A4558;",
        label: "Dark Elegance"
    },

    {
        name: "smart-style-2",
        logo: "styles/img/logo-blue.png",
        class: "btn btn-xs btn-block txt-color-darken margin-top-5",
        style: "background:#fff;",
        label: "Ultra Light"
    },

    {
        name: "smart-style-3",
        logo: "styles/img/logo-pale.png",
        class: "btn btn-xs btn-block txt-color-white margin-top-5",
        style: "background:#f78c40",
        label: "Google Skin"
    },

    {
        name: "smart-style-4",
        logo: "styles/img/logo-pale.png",
        class: "btn btn-xs btn-block txt-color-white margin-top-5",
        style: "background: #bbc0cf; border: 1px solid #59779E; color: #17273D !important;",
        label: "PixelSmash"
    },

    {
        name: "smart-style-5",
        logo: "styles/img/logo-pale.png",
        class: "btn btn-xs btn-block txt-color-white margin-top-5",
        style: "background: rgba(153, 179, 204, 0.2); border: 1px solid rgba(121, 161, 221, 0.8); color: #17273D !important;",
        label: "Glass"
    },

    {
        name: "smart-style-6",
        logo: "styles/img/logo-pale.png",
        class: "btn btn-xs btn-block txt-color-white margin-top-5",
        style: "background: #2196F3; border: 1px solid rgba(121, 161, 221, 0.8); color: #FFF !important;",
        beta: true,
        label: "MaterialDesign"
    }
];

/*
 * DEBUGGING MODE
 * debugState = true; will spit all debuging message inside browser console.
 * The colors are best displayed in chrome browser.
 */
appConfig.debugState = false;
appConfig.debugStyle = 'font-weight: bold; color: #00f;';
appConfig.debugStyle_green = 'font-weight: bold; font-style:italic; color: #46C246;';
appConfig.debugStyle_red = 'font-weight: bold; color: #ed1c24;';
appConfig.debugStyle_warning = 'background-color:yellow';
appConfig.debugStyle_success = 'background-color:green; font-weight:bold; color:#fff;';
appConfig.debugStyle_error = 'background-color:#ed1c24; font-weight:bold; color:#fff;';

appConfig.tokenKey = "token";

appConfig.PAGE_SIZE_DEFAULT = 100;
appConfig.version = 2;
appConfig.token = {"account":"xjxj"}

appConfig.AUTH_EVENTS = {
    loginSuccess: 'auth-login-success',
    loginFailed: 'auth-login-failed',
    logoutSuccess: 'auth-logout-success',
    sessionTimeout: 'auth-session-timeout',
    notAuthenticated: 'auth-not-authenticated',
    notAuthorized: 'auth-not-authorized'

}
appConfig.voice_command = true;

if (appConfig.voice_command) {

    appConfig.commands = {

        'show dashboard': function() { window.location.hash = "dashboard" },
        'show inbox': function() { window.location.hash = "inbox/" },
        'show graphs': function() { window.location.hash = "graphs/flot" },
        'show flotchart': function() { window.location.hash = "graphs/flot" },
        'show morris chart': function() { window.location.hash = "graphs/morris" },
        'show inline chart': function() { window.location.hash = "graphs/inline-charts" },
        'show dygraphs': function() { window.location.hash = "graphs/dygraphs" },
        'show tables': function() { window.location.hash = "tables/table" },
        'show data table': function() { window.location.hash = "tables/datatable" },
        'show jquery grid': function() { window.location.hash = "tables/jqgrid" },
        'show form': function() { window.location.hash = "forms/form-elements" },
        'show form layouts': function() { window.location.hash = "forms/form-templates" },
        'show form validation': function() { window.location.hash = "forms/validation" },
        'show form elements': function() { window.location.hash = "forms/bootstrap-forms" },
        'show form plugins': function() { window.location.hash = "forms/plugins" },
        'show form wizards': function() { window.location.hash = "forms/wizards" },
        'show bootstrap editor': function() { window.location.hash = "forms/other-editors" },
        'show dropzone': function() { window.location.hash = "forms/dropzone" },
        'show image cropping': function() { window.location.hash = "forms/image-editor" },
        'show general elements': function() { window.location.hash = "ui/general-elements" },
        'show buttons': function() { window.location.hash = "ui/buttons" },
        'show fontawesome': function() { window.location.hash = "ui/icons/fa" },
        'show glyph icons': function() { window.location.hash = "ui/icons/glyph" },
        'show flags': function() { window.location.hash = "ui/icons/flags" },
        'show grid': function() { window.location.hash = "ui/grid" },
        'show tree view': function() { window.location.hash = "ui/treeview" },
        'show nestable lists': function() { window.location.hash = "ui/nestable-list" },
        'show jquery U I': function() { window.location.hash = "ui/jqui" },
        'show typography': function() { window.location.hash = "ui/typography" },
        'show calendar': function() { window.location.hash = "calendar" },
        'show widgets': function() { window.location.hash = "widgets" },
        'show gallery': function() { window.location.hash = "gallery" },
        'show maps': function() { window.location.hash = "gmap-xml" },
        'go back': function() { history.back(1); },
        'scroll up': function() { $('html, body').animate({ scrollTop: 0 }, 100); },
        'scroll down': function() { $('html, body').animate({ scrollTop: $(document).height() }, 100); },
        'hide navigation': function() {
            if ($(":root").hasClass("container") && !$(":root").hasClass("menu-on-top")) {
                $('span.minifyme').trigger("click");
            } else {
                $('#hide-menu > span > a').trigger("click");
            }
        },
        'show navigation': function() {
            if ($(":root").hasClass("container") && !$(":root").hasClass("menu-on-top")) {
                $('span.minifyme').trigger("click");
            } else {
                $('#hide-menu > span > a').trigger("click");
            }
        },
        'mute': function() {
            appConfig.sound_on = false;
            $.smallBox({
                title: "MUTE",
                content: "All sounds have been muted!",
                color: "#a90329",
                timeout: 4000,
                icon: "fa fa-volume-off"
            });
        },
        'sound on': function() {
            appConfig.sound_on = true;
            $.speechApp.playConfirmation();
            $.smallBox({
                title: "UNMUTE",
                content: "All sounds have been turned on!",
                color: "#40ac2b",
                sound_file: 'voice_alert',
                timeout: 5000,
                icon: "fa fa-volume-up"
            });
        },
        'stop': function() {
            smartSpeechRecognition.abort();
            $(":root").removeClass("voice-command-active");
            $.smallBox({
                title: "VOICE COMMAND OFF",
                content: "Your voice commands has been successfully turned off. Click on the <i class='fa fa-microphone fa-lg fa-fw'></i> icon to turn it back on.",
                color: "#40ac2b",
                sound_file: 'voice_off',
                timeout: 8000,
                icon: "fa fa-microphone-slash"
            });
            if ($('#speech-btn .popover').is(':visible')) {
                $('#speech-btn .popover').fadeOut(250);
            }
        },
        'help': function() {

            $('#voiceModal').removeData('modal').modal({ remote: "app/layout/partials/voice-commands.tpl.html", show: true });
            if ($('#speech-btn .popover').is(':visible')) {
                $('#speech-btn .popover').fadeOut(250);
            }

        },
        'got it': function() {
            $('#voiceModal').modal('hide');
        },
        'logout': function() {
            $.speechApp.stop();
            window.location = $('#logout > span > a').attr("href");
        }
    };
}

/*
 * END APP.appConfig
 */