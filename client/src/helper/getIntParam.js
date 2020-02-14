//(string, string) => int
export default function getIntParam(urlNow, urlFormat) {
    try {
        const pieces = urlNow.split(urlFormat);
        return parseInt(pieces[1]);
    } catch (e) {
        console.error(e);
    }

    return -1;
}