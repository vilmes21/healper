// [{}] => [{}]
export default function makeSearchable(arr) {
    let keystr;
    for (const obj of arr) {
        if (obj.txt){
            obj.isBuilt = true
            keystr = `${obj.pinyin}${obj.txt}${obj.id}`
        } else {
            obj.txt=`${obj.namezh} ${obj.name}`
            keystr=`${obj.name}${obj.namezh}${obj.pinyin}${obj.id}`
        }
        obj.searchable = keystr
    }
    return arr;
}