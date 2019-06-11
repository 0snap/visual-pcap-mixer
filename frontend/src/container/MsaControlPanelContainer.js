import { connect } from 'react-redux'
import { createMsa } from '../actions'
import { createMsa as createMsaApiCall } from '../actions/apiActions';
import MsaControlPanel from '../representation/MsaControlPanel';

const mapStateToProps = state => {
    return {
        timeline: state.msaTimeline,
        createStatus: state.multistepattacks,
        replacements: state.replacements,
    };
}

const mapDispatchToProps = dispatch => {
    return {
        createMsa: (msa) => {
            dispatch(createMsa())
            return dispatch(createMsaApiCall(msa))
        }
    }
}

const MsaControlPanelContainer = connect(
    mapStateToProps,
    mapDispatchToProps
)(MsaControlPanel)

export default MsaControlPanelContainer