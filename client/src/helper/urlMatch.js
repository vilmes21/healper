//(string, string) => bool
export default function urlMatch(urlNow, urlFormat) {
    if (urlNow.indexOf(urlFormat) === 0) {
        const pieces = urlNow.split(urlFormat);
        if (pieces.length === 2) {
            if (pieces[1].split("/").length === 1) {
                if (!isNaN(parseInt(pieces[1]))) {
                    return true;
                }
            }
        }
    }
    return false;
}