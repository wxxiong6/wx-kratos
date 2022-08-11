import re
import sys


MODULE_REGEX = r'^[_a-zA-Z][_a-zA-Z0-9]+$'

module_name = '{{ cookiecutter.module_name }}'

if not re.match(MODULE_REGEX, module_name):
    print('ERROR: %s  不是有效的Python模块名称！' % module_name)
    # 以状态1退出，表示失败
    sys.exit(1)