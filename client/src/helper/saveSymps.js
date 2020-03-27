import config from './config'

//[[3,5], [6,7]]
export default async function saveSymps(cart) {
    try {
        const r = await fetch(config.host + config.apiUrl.createSymps, {
            method: 'POST',
            body: JSON.stringify(cart)
        });
        const d = await r.json();
        return d;
    } catch (e) {
        console.error('saveSymps Error:', e);
    }

    return false;
}
