import React, { Component } from 'react';
import './App.css';
import AttackListContainer from './container/AttackListContainer';
import TraceFileListContainer from './container/TraceFileListContainer';
import ReplacementListContainer from './container/ReplacementListContainer';
import ReplacementFormContainer from './container/ReplacementFormContainer';
import MsaTimelineContainer from './container/MsaTimelineContainer';
import MsaControlPanelContainer from './container/MsaControlPanelContainer';
import MsaListContainer from './container/MsaListContainer';

class App extends Component {
    render() {
        return (
            <div className="App">
                <div className="Panel">
                    <AttackListContainer />
                    <TraceFileListContainer />
                </div>
                <div className="Panel">
                    <MsaTimelineContainer />
                </div>
                <div className="Panel">
                <div className="HalfPanel">
                    <ReplacementFormContainer />
                    <ReplacementListContainer />
                </div>
                <div className="HalfPanel">
                    <MsaControlPanelContainer />
                    <MsaListContainer />
                </div>
                </div>
            </div>
        );
    }
}

export default App;
