import subprocess
import datetime

def format_go_files():
    result = subprocess.run(
        ['find', '.', '-name', '*.go', '-not', '-path', './vendor/*', '-not', '-path', './local_pgdata/*'],
        capture_output=True,
        text=True
    )
    go_files = result.stdout.splitlines()

    for f in go_files:
        subprocess.call(['go', 'fmt', f])

subprocess.call(['sleep', '2'])

format_go_files()

status = subprocess.check_output(['git', 'status', '--porcelain']).decode().strip()
if not status:
    print("No changes No commit needed.")
    exit(0)

subprocess.check_call(['git', 'add', '.'])
msg = f'fmt: format go files {datetime.datetime.now().strftime("%Y-%m-%d %H:%M")}'
subprocess.check_call(['git', 'commit', '-m', msg])
subprocess.check_call(['git', 'push'])

print("✓ Commit and push completed.")
