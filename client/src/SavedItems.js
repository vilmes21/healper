import React from "react";
import SavedItem from "./SavedItem";

class SavedItems extends React.Component {
    render() {
        //items-shape== [{id, txt},{}]
        const {items, ownerId} = this.props;

        if (items.length === 0) {
            return null
        }

        return <div>
            {items.map(y => {
                return <SavedItem key={`${ownerId}-${y.organId}-${y.symptomId}`} data={y}/>
            })}
        </div>
    }
}

export default SavedItems;