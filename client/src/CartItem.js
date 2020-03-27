import React from "react";

class CartItem extends React.Component {
    render() {
        //data-shape=[{},{}]
        const {data,del} = this.props;
        let txt;
        if (Array.isArray(data)){
            const txtArr = data.map(x => x.namezh);
             txt = txtArr.join("")
        } else {
            txt = data.txt
        }
        
        return <div className="fontsize5 inlineBloc cartitem">
            {txt}
            <span 
            onClick={del}
            className="del1 whitecolor pointercursor">
            &#x2715;
            </span>
        </div>
    }
}

export default CartItem;