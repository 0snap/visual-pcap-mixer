import { connect } from 'react-redux'
import MsaList from '../representation/MsaList';
import { loadMsaToTimeline } from '../actions'


const mapStateToProps = state => {
    return {
        multistepattacks: state.multistepattacks.knownAttacks,
    };
}

const mapDispatchToProps = dispatch => {
    return {
        loadMsaToTimeline: msa => {
            dispatch(loadMsaToTimeline(msa));
        }
    }
}

const MsaListContainer = connect(
    mapStateToProps,
    mapDispatchToProps
)(MsaList)

export default MsaListContainer