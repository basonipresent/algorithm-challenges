var __extends = (this && this.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (Object.prototype.hasOwnProperty.call(b, p)) d[p] = b[p]; };
        return extendStatics(d, b);
    };
    return function (d, b) {
        if (typeof b !== "function" && b !== null)
            throw new TypeError("Class extends value " + String(b) + " is not a constructor or null");
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();
var Dialog = /** @class */ (function () {
    function Dialog() {
    }
    Dialog.prototype.doRender = function () {
        var button = this.createButton();
        return "Dialog : render button " + button.render();
    };
    return Dialog;
}());
var WindowsDialog = /** @class */ (function (_super) {
    __extends(WindowsDialog, _super);
    function WindowsDialog() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    WindowsDialog.prototype.createButton = function () {
        return new WindowsButton();
    };
    return WindowsDialog;
}(Dialog));
var WebDialog = /** @class */ (function (_super) {
    __extends(WebDialog, _super);
    function WebDialog() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    WebDialog.prototype.createButton = function () {
        return new HTMLButton();
    };
    return WebDialog;
}(Dialog));
var WindowsButton = /** @class */ (function () {
    function WindowsButton() {
    }
    WindowsButton.prototype.render = function () {
        return "Result Windows Button";
    };
    return WindowsButton;
}());
var HTMLButton = /** @class */ (function () {
    function HTMLButton() {
    }
    HTMLButton.prototype.render = function () {
        return "Result HTML Button";
    };
    return HTMLButton;
}());
var main = function (dialog) {
    console.log(dialog.doRender());
};
main(new WindowsDialog());
main(new WebDialog());
