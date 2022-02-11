abstract class Dialog {
  public abstract createButton() : Button;
  public doRender() : string {
    const button = this.createButton();
    return `Dialog : render button ${button.render()}`;
  }
}

class WindowsDialog extends Dialog {
  public createButton() : Button {
    return new WindowsButton();
  }
}

class WebDialog extends Dialog {
  public createButton() : Button {
    return new HTMLButton();
  }
}

interface Button {
  render() : string;
}

class WindowsButton implements Button {
  public render() : string {
    return `Result Windows Button`;
  }
}

class HTMLButton implements Button {
  public render() : string {
    return `Result HTML Button`;
  }
}

const main = (dialog : Dialog) => {
  console.log(dialog.doRender());
}

main(new WindowsDialog());
main(new WebDialog());