import React from "react";
import Link2 from "./Link2";
import config from "./helper/config";

class Home extends React.Component {
  render(){
    return <div>
     Home

     <Link2 to={config.urls.newDisease}>
      Add disease
     </Link2>
    </div>
  }
}

export default Home;