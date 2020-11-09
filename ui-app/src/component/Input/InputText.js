import TextField from '@material-ui/core/TextField';

export function InputText(props) {
  const { errorMsg, ...rest  } = props
  let error = false, helperText ="";
  if (errorMsg) {
    error = true;
    helperText = errorMsg;
  }
  return (<TextField
    error={error}
    variant="outlined"
    margin="normal"
    fullWidth
    helperText={helperText}
    {...rest}
  />)
}