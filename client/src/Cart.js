import React from "react";
import CartItem from "./CartItem";

class Cart extends React.Component {
    render() {

      //objArr-shape==[[{},{}], [{},{}], {}] is not consistent
        const {objArr,del,save,itemTypeStr} = this.props;
        console.log(`Cart compo, objArr:`, objArr)

        if (objArr.length === 0){
          return null
        }

        return <div>
            {objArr.map(y => {
              let _key=`${y.id}-${y.organId}-${y.symptomId}`;

              return <CartItem
              del={del(y)}
             key={_key} 
             data={y}/>})}

            <button className="bt" onClick={save}>Save {itemTypeStr}</button>
        </div>
    }
}

export default Cart;