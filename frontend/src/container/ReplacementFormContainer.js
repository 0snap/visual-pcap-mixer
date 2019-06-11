import { connect } from 'react-redux'
import { addReplacement } from '../actions'
import ReplacementForm from '../representation/ReplacementForm';

const mapStateToProps = state => {
    return {};
}

const mapDispatchToProps = dispatch => {
    return {
        addReplacement: (ipA, ipB) => {
            dispatch(addReplacement(ipA, ipB))
        }
    }
}

const ReplacementFormContainer = connect(
    mapStateToProps,
    mapDispatchToProps
)(ReplacementForm)

export default ReplacementFormContainer