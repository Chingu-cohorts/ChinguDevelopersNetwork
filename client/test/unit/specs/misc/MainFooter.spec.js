import { mount } from 'avoriaz'
import MainFooter from '@/components/misc/MainFooter'

describe('MainFooter.vue', () => {
  it('should render the main footer component', () => {
    const wrapper = mount(MainFooter)
    expect(wrapper.is('footer')).to.equal(true)
    expect(wrapper.hasClass('footer')).to.equal(true)
  })

  it('should have a computed property to calculate the year', () => {
    const wrapper = mount(MainFooter)
    expect(typeof wrapper.computed().currentYear).to.equal('function')
  })
})
