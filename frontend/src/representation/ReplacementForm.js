import React, { Component } from 'react';
import PropTypes from 'prop-types';
import './Form.css';

class ReplacementForm extends Component {

    constructor(props) {
        super(props);

        this.state = {
            source: '',
            target: '',
        };
        this.handleSourceChange = this.handleSourceChange.bind(this);
        this.handleTargetChange = this.handleTargetChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleSourceChange(event) {
        this.setState({source: event.target.value});
    }

    handleTargetChange(event) {
        this.setState({target: event.target.value});
    }

    handleSubmit(event) {
        event.preventDefault();
        const {source, target} = this.state;
        if (!source || !target) {
            return;
        }
        source.split(',').map(src_ip => this.props.addReplacement(src_ip, this.state.target));
        this.setState({
            source: '',
            target: ''
        });
    }

    render() {
        return (
            <div className='Form'>
                <h3>IP addresses</h3>
                <p>The entered IP addresses will be replaced in each PCAP of the timeline.</p>
                <form onSubmit={this.handleSubmit}>
                <div className='FormInput'>
                    <label>
                        Sources (single IP or comma delimited list):
                        <input type='text' value={this.state.source} onChange={this.handleSourceChange} />
                    </label>
                </div>
                <div className='FormInput'>
                    <label>
                        Target (single IP):
                        <input type='text' value={this.state.target} onChange={this.handleTargetChange} />
                    </label>
                </div>
                <div className='FormInput'>
                    <input className='submit' type='submit' value='Add Replacement' />
                </div>
                </form>
            </div>
        );
    }
}

ReplacementForm.propTypes = {
    addReplacement: PropTypes.func.isRequired,
}

export default ReplacementForm;