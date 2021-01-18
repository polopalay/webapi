function includeHTML() {
    let sections;
    sections = document.getElementsByTagName("SECTION");
    for (let i = 0; i < sections.length; i++) {
        let elmnt = sections[i];
        let file = elmnt.getAttribute("include-html");
        if (file != null) {
            fetch(file).then(rs => {
                switch (rs.status) {
                    case 200:
                        rs.text().then(text => { elmnt.innerHTML = text })
                        break;
                    default:
                        break;
                }
                elmnt.removeAttribute("include-html");
            })
        }
    }
}

includeHTML()