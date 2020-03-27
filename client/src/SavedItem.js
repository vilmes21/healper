import React from "react";

class SavedItem extends React.Component {
    render() {
        //data-shape=[{},{}]
        const {data} = this.props;
        
        return <div className="fontsize5 inlineBloc saveditem">
            {data.txt}
        </div>
    }
}

export default SavedItem;