import React, { Component } from 'react';
import PropTypes from 'prop-types';
import classNames from 'classnames'
import styles from './Form.css';

class MsaControlPanl extends Component {

    constructor(props) {
        super(props);

        this.state = {
            msaName: '',
        };
        this.handleNameChange = this.handleNameChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleSubmit(event) {
        event.preventDefault();
        const {timeline, replacements, createMsa} = this.props;
        const name = this.state.msaName;
        if (!name) {
            return;
        }
        const msaCreatePayload = {
            timeline,
            replacements,
            name,
        }

        createMsa(msaCreatePayload);
        this.setState({msaName: ''});
    }

    handleNameChange(event) {
        this.setState({msaName: event.target.value});
    }

    render() {
        const { createStatus } = this.props;
        const cx = classNames.bind(styles);
        const btnState = cx({
            ...createStatus,
            submit: true,
        });

        return (
            <div className='Form'>
                <h3>Create Multi Step Attack</h3>
                <p>Create a Multi Step Attack from the current timeline. This will synchronize all packet timestamps to appear to have been recorded in sequence and all IP address replacements will be applied. A folder will be generated with the name of the MSA.</p>
                <form onSubmit={this.handleSubmit}>
                <div className='FormInput'>
                    <label>
                        MultiStepAttack Name:
                        <input type='text' value={this.state.msaName} onChange={this.handleNameChange} />
                    </label>
                </div>
                <div className='FormInput'>
                    <input type='submit' className={btnState} value='Create MSA' disabled={createStatus.waiting}/>
                </div>
                </form>
            </div>
        );
    }
}

MsaControlPanl.propTypes = {
    timeline: PropTypes.array.isRequired,
    replacements: PropTypes.array.isRequired,
    createStatus: PropTypes.object.isRequired,
    createMsa: PropTypes.func.isRequired,
}

export default MsaControlPanl;