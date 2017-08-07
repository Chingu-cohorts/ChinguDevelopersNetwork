import { mount } from 'avoriaz'
import Privacy from '@/components/misc/Privacy'

describe('Privacy.vue', () => {
  it('should render the privacy component', () => {
    const wrapper = mount(Privacy)
    expect(wrapper.is('section')).to.equal(true)
  })
})
