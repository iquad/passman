import React, {Component} from "react";
import {Provider} from "react-redux";
import {createStore} from "redux";
import {reducer} from "../reducer/reducer.js";


class App extends Component {
    constructor(props) {
        super(props);
        this.store = createStore(reducer);
    }

    render() {
        return (
            <Provider store={this.store}>
                <div className="App">
                    <div className={"container"}>
                        <div className={"title"}>
                            {"To33do"}
                        </div>
                    </div>
                </div>
            </Provider>
        );
    }
}

export default App;