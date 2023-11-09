const useTap = ({target, onDown, onUp, onSuccess, onCancel}) => {
  target.onmousedown = e => {
    onDown?.(e)
    const preventSelect = e => e.preventDefault()
    const handleUp = e => {
      onUp?.(e);
      if (Object.is(e.target, target)) {
        onSuccess?.(e)
      } else {
        onCancel?.(e)
      }
      document.removeEventListener('mouseup', handleUp)
      document.removeEventListener('selectstart', preventSelect)
    }
    document.addEventListener('selectstart', preventSelect)
    document.addEventListener('mouseup', handleUp)
  }
}