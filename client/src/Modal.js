import React from "react";


class Modal extends React.Component {
    render() {
        const {children, setModal}=this.props;
        
        return <div className="absPos blackbg fullWidth fullHeight modal1 ">
            <div onClick={setModal(false)}
            className="whitecolor">Close</div>
            <div className="whitebg modal2">
                {children}
            </div>
        </div>
    }
}

export default Modal;